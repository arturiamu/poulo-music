package migu

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"poulo-music/config"
	"poulo-music/httpp"
	"poulo-music/models"
	"poulo-music/platform"
)

var _ platform.Platform = (*Migu)(nil)

// Migu
// 请求方式：get/post
// 请求地址：http://www.ilingku.com/api/migu/v2
// 返回格式：json
// 返回参数：
// 名称	类型	说明
// code	string	返回的状态码
// msg	string	返回错误提示！
type Migu struct {
	log *logrus.Logger
	cfg *config.Config
}

func NewMigu(log *logrus.Logger, cfg *config.Config) *Migu {
	return &Migu{log: log, cfg: cfg}
}

func (m *Migu) GetSearch(ctx context.Context, param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	miguSearchResp, err := m.search(param.Keyword, param.Pagesize)
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	for _, v := range miguSearchResp.Data {
		data = append(data, models.GetSearchResp{
			ID:         v.Id,
			Platform:   models.PlatformMigu,
			Identifier: v.Id,
			Title:      v.Name,   //一路向北
			Name:       v.Singer, //周杰伦
			Cover:      v.Pic,
		})
	}
	return
}

func (m *Migu) GetHotContent(ctx context.Context, param models.GetHotContentParam) (data []models.GetHotContentResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *Migu) GetMusic(ctx context.Context, param models.GetMusicParam) (data models.Music, err error) {
	miguPlayInfoResp, err := m.info(param.Identifier)
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	lrc := fmt.Sprintf("http://www.ilingku.com/api/migu/v2?act=lrcgc&cid=%s", param.Identifier)

	data = models.Music{
		Platform:   models.PlatformMigu,
		Identifier: param.Identifier,
		Name:       miguPlayInfoResp.Name,
		Artist:     miguPlayInfoResp.Singer,
		Url:        miguPlayInfoResp.Url,
		Cover:      miguPlayInfoResp.Pic,
		Lrc:        lrc,
	}

	return
}

// 搜索请求示例：http://www.ilingku.com/api/migu/v2?act=search&total=5&wd=%E6%AC%A2%E5%AD%90
// 搜索请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|info=歌曲信息|lrcgc=歌词|music=音乐url）
// wd	是	string	需要搜索的关键词
// total	否	string	输出数据条数（默认10条）
func (m *Migu) search(keyword string, total int64) (data models.MiguSearchResp, err error) {
	if total <= 0 {
		total = 10
	}

	var searchUrl = "http://www.ilingku.com/api/migu/v2"
	var params = map[string]string{"act": "search", "wd": keyword, "total": cast.ToString(total)}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌曲请求示例：http://www.ilingku.com/api/migu/v2?act=info&cid=69001000008
// 歌曲信息请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的咪咕音乐id
func (m *Migu) info(cid string) (data models.MiguInfoResp, err error) {
	var searchUrl = "http://www.ilingku.com/api/migu/v2"
	var params = map[string]string{"act": "info", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		m.log.Error(err.Error())
	}
	return
}

// 歌词请求示例：http://www.ilingku.com/api/migu/v2?act=lrcgc&cid=69001000008
// 歌词请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的网易云音乐id
func (m *Migu) lrcgc(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/migu/v2"
	var params = map[string]string{"act": "lrcgc", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	fmt.Println(string(bytes))

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		m.log.Error(err.Error())
	}
	return
}

// 音乐请求示例：http://www.ilingku.com/api/migu/v2?act=music&format=url&cid=69001000008
func (m *Migu) music(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/migu/v2"
	var params = map[string]string{"act": "music", "cid": cid, "format": "url"}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		m.log.Error(err.Error())
		return
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		m.log.Error(err.Error())
	}
	return
}
