package models

type QQSearchData struct {
	Id        int64  `json:"id"`
	Mid       string `json:"mid"`
	Name      string `json:"name"`
	Albummid  string `json:"albummid"`
	Albumname string `json:"albumname"`
	Singer    string `json:"singer"`
	Singerid  string `json:"singerid"`
	Pic       string `json:"pic"`
	Time      string `json:"time"`
	Info      string `json:"info"`
	Lrcgc     string `json:"lrcgc"`
}

type QQSearchResp struct {
	Num  int64          `json:"num"`
	List []QQSearchData `json:"list"`
}

type QQPlayInfoResp struct {
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Name      string `json:"name"`
	Singer    string `json:"singer"`
	Singerid  int64  `json:"singerid"`
	Singermid string `json:"singermid"`
	Pic       string `json:"pic"`
	Url       string `json:"url"`
	Lrcgc     string `json:"lrcgc"`
}
