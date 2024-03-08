package models

type KuwoInfoResp struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Name      string      `json:"name"`
	Singer    string      `json:"singer"`
	Singerid  string      `json:"singerid"`
	Singerpic string      `json:"singerpic"`
	Albumid   string      `json:"albumid"`
	Albumname string      `json:"albumname"`
	Payvip    string      `json:"payvip"`
	Time      string      `json:"time"`
	Pic       string      `json:"pic"`
	Url       interface{} `json:"url"`
	Lrc       string      `json:"lrc"`
}
