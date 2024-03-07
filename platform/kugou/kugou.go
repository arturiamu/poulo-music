package netease

import (
	"encoding/json"
	"github.com/spf13/cast"
	"poulo-music/httpp"
	"poulo-music/models"
)

// Kugou
// 请求方式：get/post
// 请求地址：http://www.ilingku.com/api/kugou/v2
// 返回格式：json
// 返回参数：
// 名称	类型	说明
// code	string	返回的状态码
// msg	string	返回错误提示！
type Kugou struct{}

// 搜索请求示例：http://www.ilingku.com/api/kugou/v2?act=search&total=5&wd=%E6%AC%A2%E5%AD%90
// 搜索请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|info=歌曲信息|lrcgc=歌词|music=音乐url）
// wd	是	string	需要搜索的关键词
// total	否	string	输出数据条数（默认10条）
func (k *Kugou) search(keyword string, total int64) (data models.KugouSearchResp, err error) {
	if total <= 0 {
		total = 10
	}

	var searchUrl = "http://www.ilingku.com/api/kugou/v2"
	var params = map[string]string{"act": "search", "wd": keyword, "total": cast.ToString(total)}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌曲请求示例：http://www.ilingku.com/api/kugou/v2?act=info&cid=772cf033030e7063f8e34aa083787e58
// 歌曲信息请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的网易云音乐id
func (k *Kugou) info(cid string) (data models.KugouInfoResp, err error) {
	var searchUrl = "http://www.ilingku.com/api/kugou/v2"
	var params = map[string]string{"act": "info", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 歌词请求示例：http://www.ilingku.com/api/kugou/v2?act=lrcgc&cid=772cf033030e7063f8e34aa083787e58
// 歌词请求参数：
// 名称	必填	类型	说明
// act	是	string	需要解析的类型（search=搜索|playinfo=歌曲信息|lrcgc=歌词|music=音乐url）
// cid	是	string	需要查询的网易云音乐id
func (k *Kugou) lrcgc(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/kugou/v2"
	var params = map[string]string{"act": "lrcgc", "cid": cid}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}

// 音乐请求示例：http://www.ilingku.com/api/kugou/v2?act=music&format=url&cid=772cf033030e7063f8e34aa083787e58
func (k *Kugou) music(cid string) (data any, err error) {
	var searchUrl = "http://www.ilingku.com/api/kugou/v2"
	var params = map[string]string{"act": "music", "cid": cid, "format": "url"}

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).AddParams(params).Do()
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &data)
	return
}
