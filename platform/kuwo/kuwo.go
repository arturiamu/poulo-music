package kuwo

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cast"
	"poulo-music/httpp"
	"poulo-music/models"
)

// Kuwo
// 提示：酷我音乐歌曲url已支持320高音质VIP数据
// 请求方式：get/post
// 请求地址：http://www.ilingku.com/api/kuwo/v2
// 返回格式：json
// 返回参数：
// 名称	类型	说明
// code	string	返回的状态码
// msg	string	返回错误提示！
type Kuwo struct{}

// 搜索请求示例：http://www.ilingku.com/api/kuwo/v2?act=search&total=5&wd=%E6%AC%A2%E5%AD%90
// 搜索请求参数：
// act	是	string	需要解析的类型（search=搜索|info=歌曲信息|lrcgc=歌词|music=音乐url）
// wd	是	string	需要搜索的关键词
// total	否	string	输出数据条数（默认10条）
func (k *Kuwo) search(keyword string, total int64) (data any, err error) {
	return nil, errors.New("not implemented")

	if total <= 0 {
		total = 10
	}

	var searchUrl = "http://www.ilingku.com/api/kuwo/v2"
	var params = map[string]string{"act": "search", "wd": keyword, "total": cast.ToString(total)}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌曲请求示例：http://www.ilingku.com/api/kuwo/v2?act=info&cid=274850577
// 歌曲信息请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的酷我音乐id
func (k *Kuwo) info(cid string) (data models.KuwoInfoResp, err error) {
	var searchUrl = "http://www.ilingku.com/api/kuwo/v2"
	var params = map[string]string{"act": "info", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌词请求示例：http://www.ilingku.com/api/kuwo/v2?act=lrcgc&cid=274850577
// 歌词请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的酷我音乐id
func (k *Kuwo) lrcgc(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/kuwo/v2"
	var params = map[string]string{"act": "lrcgc", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 音乐请求示例：http://www.ilingku.com/api/kuwo/v2?act=music&format=url&cid=274850577
func (k *Kuwo) music(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/kuwo/v2"
	var params = map[string]string{"act": "music", "cid": cid, "format": "url"}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 视频请求示例：http://www.ilingku.com/api/kuwo/v2?act=video&format=url&cid=82182179
func (k *Kuwo) video(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/kuwo/v2"
	var params = map[string]string{"act": "video", "cid": cid, "format": "url"}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}
