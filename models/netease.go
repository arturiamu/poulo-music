package models

type NeteaseSearchData struct {
	Id     int    `json:"id"`
	Mvid   int    `json:"mvid"`
	Name   string `json:"name"`
	Singer string `json:"singer"`
	Album  string `json:"album"`
	Title  string `json:"title"`
	Lrcgc  string `json:"lrcgc"`
	Info   string `json:"info"`
	Url    string `json:"url"`
}

type NeteaseSearchResp struct {
	Num   int                 `json:"num"`
	Total int                 `json:"total"`
	List  []NeteaseSearchData `json:"list"`
}

type NeteasePlayInfoResp struct {
	Code             int    `json:"code"`
	Msg              string `json:"msg"`
	Title            string `json:"title"`
	Name             string `json:"name"`
	Singer           string `json:"singer"`
	Singerid         int    `json:"singerid"`
	Singerpic        string `json:"singerpic"`
	Singersongcount  int    `json:"singersongcount"`
	Albumname        string `json:"albumname"`
	Albumid          int    `json:"albumid"`
	Singeralbumcount int    `json:"singeralbumcount"`
	Url              string `json:"url"`
	Pic              string `json:"pic"`
	Lrc              string `json:"lrc"`
}
