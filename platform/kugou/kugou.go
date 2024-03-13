package kugou

import (
	"encoding/json"
	"fmt"
	"poulo-music/httpp"
	"poulo-music/models"
	"strings"
)

type Kugou struct{}

/**
 * POST /api/music/search
 * @tags 酷狗音乐
 * @summary 音乐搜索
 * @description 酷狗音乐搜索
 * @param {string}  keyword.params.required  -  搜索关键字
 */
func (k *Kugou) search(keyword string, page int64, pagesize int64) (data models.KugouSearchResp, err error) {
	if page <= 0 {
		page = 1
	}

	if pagesize <= 0 {
		pagesize = 20
	}

	searchUrl := "http://mobilecdn.kugou.com/api/v3/search/song?format=json&keyword=%s&page=%d&pagesize=%d&showtype=1"
	searchUrl = fmt.Sprintf(searchUrl, keyword, page, pagesize)

	bytes, err := httpp.NewHttp().SetUrl(searchUrl).Do()
	str := string(bytes)
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")

	err = json.Unmarshal([]byte(str), &data)
	return
}

/**
 * POST /api/music/song-info
 * @tags 酷狗音乐
 * @summary  音乐详情信息
 * @description  音乐详情信息
 */
func (k *Kugou) info(hash string) (data models.KugouInfoResp, err error) {
	var infoUrl = "http://m.kugou.com/app/i/getSongInfo.php?cmd=playInfo&hash=%s"
	infoUrl = fmt.Sprintf(infoUrl, hash)
	bytes, err := httpp.NewHttp().SetUrl(infoUrl).Do()
	if err != nil {
		return
	}

	str := string(bytes)
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")

	fmt.Println(str)

	err = json.Unmarshal([]byte(str), &data)
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
