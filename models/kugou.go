package models

type KugouSearchData struct {
	Id       string `json:"id"`
	Mvid     string `json:"mvid"`
	Title    string `json:"title"`
	Name     string `json:"name"`
	Singer   string `json:"singer"`
	Filesize string `json:"filesize"`
	Lrcgc    string `json:"lrcgc"`
	Info     string `json:"info"`
	Url      string `json:"url"`
}

type KugouSearchResp struct {
	Code  int               `json:"code"`
	Msg   string            `json:"msg"`
	Total int               `json:"total"`
	Data  []KugouSearchData `json:"data"`
}

type KugouInfoResp struct {
	Code            int    `json:"code"`
	Msg             string `json:"msg"`
	Id              string `json:"id"`
	Title           string `json:"title"`
	Name            string `json:"name"`
	Singer          string `json:"singer"`
	Singerid        int    `json:"singerid"`
	Singerpic       string `json:"singerpic"`
	Singersongcount int    `json:"singersongcount"`
	Pic             string `json:"pic"`
	Album           string `json:"album"`
	Albumid         int    `json:"albumid"`
	AlbumAudioId    int    `json:"album_audio_id"`
	Time            int    `json:"time"`
	Addtime         string `json:"addtime"`
	Url             string `json:"url"`
	Lrc             string `json:"lrc"`
}
