package models

type MiguSearchData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Singerid string `json:"singerid"`
	Singer   string `json:"singer"`
	Pic      string `json:"pic"`
	Lrcgc    string `json:"lrcgc"`
	Info     string `json:"info"`
	Url      string `json:"url"`
}

type MiguSearchResp struct {
	Code      int              `json:"code"`
	Msg       string           `json:"msg"`
	Count     int              `json:"count"`
	Total     string           `json:"total"`
	Pagecount int              `json:"pagecount"`
	Data      []MiguSearchData `json:"data"`
}

type MiguInfoResp struct {
	Name      string `json:"name"`
	Singer    string `json:"singer"`
	Singerid  string `json:"singerid"`
	Album     string `json:"album"`
	Albumid   string `json:"albumid"`
	PlayTime  string `json:"play_time"`
	Hits      int    `json:"hits"`
	HotNum    string `json:"hot_num"`
	Singerpic string `json:"singerpic"`
	Pic       string `json:"pic"`
	Img       string `json:"img"`
	Url       string `json:"url"`
	Flacsize  string `json:"flacsize"`
	Addtime   bool   `json:"addtime"`
}
