package qq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"path/filepath"
	"poulo-music/config"
	"poulo-music/httpp"
	"poulo-music/models"
	"poulo-music/platform"
	"strings"
)

const (
	Sep = string(filepath.Separator)
)

var _ platform.Platform = (*QQ)(nil)

// QQ
// 提示：QQ音乐歌曲url已支持VIP数据
// 请求方式：get/post
// 请求地址：http://www.ilingku.com/api/qqmusic/v2
// 返回格式：json
// 返回参数：
// code	string	返回的状态码
// msg	string	返回错误提示！
type QQ struct {
	log *logrus.Logger
	cfg *config.Config
}

func NewQQ(log *logrus.Logger, cfg *config.Config) *QQ {
	return &QQ{log: log, cfg: cfg}
}

func (q *QQ) GetSearch(ctx context.Context, param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	param.Keyword = strings.ReplaceAll(param.Keyword, " ", "")
	qqSearchResp, err := q.search(param.Keyword, param.Pagesize)
	if err != nil {
		q.log.Error(err.Error())
		return
	}

	for _, v := range qqSearchResp.List {
		data = append(data, models.GetSearchResp{
			ID:         cast.ToString(v.Id),
			Platform:   models.PlatformQQ,
			Identifier: v.Mid,
			Title:      v.Name,      //一路向北
			Name:       v.Singer,    //周杰伦
			Describe:   v.Albumname, //J III MP3 Player
			Cover:      v.Pic,
		})
	}
	return
}

func (q *QQ) GetHotContent(ctx context.Context, param models.GetHotContentParam) (data []models.GetHotContentResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (q *QQ) GetMusic(ctx context.Context, param models.GetMusicParam) (data models.Music, err error) {
	qqPlayInfoResp, err := q.playInfo(param.Identifier)
	if err != nil {
		q.log.Error(err.Error())
		return
	}

	data = models.Music{
		Platform:   models.PlatformQQ,
		Identifier: param.Identifier,
		Name:       qqPlayInfoResp.Name,
		Artist:     qqPlayInfoResp.Singer,
		Url:        qqPlayInfoResp.Url,
		Cover:      qqPlayInfoResp.Pic,
		Lrc:        qqPlayInfoResp.Lrcgc,
	}

	if param.CacheLrc {
		filename := "qq_" + param.Identifier + ".lrc"
		path := q.cfg.BaseDir + Sep + "cache" + Sep + filename
		err = httpp.NewHttp().SetUrl(qqPlayInfoResp.Lrcgc).SaveFile(path)
		if err != nil {
			q.log.Error(err.Error())
			return data, nil
		}
		data.Lrc = fmt.Sprintf("http://localhost:%d/%s", q.cfg.FileServerPort, filename)
	}
	return
}

// 搜索请求示例：http://www.ilingku.com/api/qqmusic/v2?act=search&total=5&wd=%E5%91%A8%E6%9D%B0%E4%BC%A6
// 搜索请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词）
// wd	是	string	需要搜索的关键词
// total	否	string	输出数据条数（默认10条）
func (q *QQ) search(keyword string, total int64) (data models.QQSearchResp, err error) {
	if total <= 0 {
		total = 10
	}

	var searchUrl = "http://www.ilingku.com/api/qqmusic/v2"
	var params = map[string]string{"act": "search", "wd": keyword, "total": cast.ToString(total)}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		q.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌曲请求示例：http://www.ilingku.com/api/qqmusic/v2?act=playinfo&cid=001zMQr71F1Qo8
// 歌曲信息请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词）
// cid	是	string	需要查询的QQ音乐mid
func (q *QQ) playInfo(cid string) (data models.QQPlayInfoResp, err error) {
	var searchUrl = "http://www.ilingku.com/api/qqmusic/v2"
	var params = map[string]string{"act": "playinfo", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		q.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌词请求示例：http://www.ilingku.com/api/qqmusic/v2?act=lrcgc&cid=001zMQr71F1Qo8
// 歌词请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词）
// cid	是	string	需要查询的QQ音乐mid
func (q *QQ) lrcgc(cid string) (data models.QQPlayInfoResp, err error) {
	var searchUrl = "http://www.ilingku.com/api/qqmusic/v2"
	var params = map[string]string{"act": "lrcgc", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		q.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}
