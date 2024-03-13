package models

type KugouSearchData struct {
	Timestamp       int    `json:"timestamp"`
	Tab             string `json:"tab"`
	Forcecorrection int    `json:"forcecorrection"`
	Correctiontype  int    `json:"correctiontype"`
	Total           int    `json:"total"`
	Istag           int    `json:"istag"`
	Allowerr        int    `json:"allowerr"`
	Info            []struct {
		Hash              string `json:"hash"`
		Sqfilesize        int    `json:"sqfilesize"`
		Sourceid          int    `json:"sourceid"`
		PayTypeSq         int    `json:"pay_type_sq"`
		Bitrate           int    `json:"bitrate"`
		Ownercount        int    `json:"ownercount"`
		PkgPriceSq        int    `json:"pkg_price_sq"`
		Songname          string `json:"songname"`
		AlbumName         string `json:"album_name"`
		SongnameOriginal  string `json:"songname_original"`
		Accompany         int    `json:"Accompany"`
		Sqhash            string `json:"sqhash"`
		FailProcess       int    `json:"fail_process"`
		PayType           int    `json:"pay_type"`
		RpType            string `json:"rp_type"`
		AlbumId           string `json:"album_id"`
		OthernameOriginal string `json:"othername_original"`
		Mvhash            string `json:"mvhash"`
		Extname           string `json:"extname"`
		Group             []struct {
			Hash              string `json:"hash"`
			Sqfilesize        int    `json:"sqfilesize"`
			Sourceid          int    `json:"sourceid"`
			PayTypeSq         int    `json:"pay_type_sq"`
			Bitrate           int    `json:"bitrate"`
			Ownercount        int    `json:"ownercount"`
			PkgPriceSq        int    `json:"pkg_price_sq"`
			Songname          string `json:"songname"`
			AlbumName         string `json:"album_name"`
			SongnameOriginal  string `json:"songname_original"`
			Accompany         int    `json:"Accompany"`
			Sqhash            string `json:"sqhash"`
			FailProcess       int    `json:"fail_process"`
			PayType           int    `json:"pay_type"`
			RpType            string `json:"rp_type"`
			AlbumId           string `json:"album_id"`
			OthernameOriginal string `json:"othername_original"`
			Mvhash            string `json:"mvhash"`
			Extname           string `json:"extname"`
			Price320          int    `json:"price_320"`
			Hash1             string `json:"320hash"`
			Topic             string `json:"topic"`
			Othername         string `json:"othername"`
			Isnew             int    `json:"isnew"`
			FoldType          int    `json:"fold_type"`
			OldCpy            int    `json:"old_cpy"`
			Srctype           int    `json:"srctype"`
			Singername        string `json:"singername"`
			AlbumAudioId      int    `json:"album_audio_id"`
			Duration          int    `json:"duration"`
			Filesize          int    `json:"320filesize"`
			PkgPrice320       int    `json:"pkg_price_320"`
			AudioId           int    `json:"audio_id"`
			Feetype           int    `json:"feetype"`
			Price             int    `json:"price"`
			Filename          string `json:"filename"`
			Source            string `json:"source"`
			PriceSq           int    `json:"price_sq"`
			FailProcess320    int    `json:"fail_process_320"`
			TransParam        struct {
				CpyLevel int `json:"cpy_level,omitempty"`
				CpyGrade int `json:"cpy_grade,omitempty"`
				Classmap struct {
					Attr0 int `json:"attr0"`
				} `json:"classmap"`
				PayBlockTpl int `json:"pay_block_tpl"`
				Qualitymap  struct {
					Attr0 int `json:"attr0"`
				} `json:"qualitymap"`
				AppidBlock     string `json:"appid_block,omitempty"`
				HashMultitrack string `json:"hash_multitrack,omitempty"`
				Language       string `json:"language"`
				Cid            int    `json:"cid"`
				CpyAttr0       int    `json:"cpy_attr0"`
				Ipmap          struct {
					Attr0 int `json:"attr0"`
				} `json:"ipmap"`
				HashOffset struct {
					ClipHash   string `json:"clip_hash"`
					StartByte  int    `json:"start_byte"`
					EndMs      int    `json:"end_ms"`
					EndByte    int    `json:"end_byte"`
					FileType   int    `json:"file_type"`
					StartMs    int    `json:"start_ms"`
					OffsetHash string `json:"offset_hash"`
				} `json:"hash_offset,omitempty"`
				MusicpackAdvance int `json:"musicpack_advance"`
				Display          int `json:"display"`
				DisplayRate      int `json:"display_rate"`
			} `json:"trans_param"`
			PkgPrice      int    `json:"pkg_price"`
			PayType320    int    `json:"pay_type_320"`
			TopicUrl      string `json:"topic_url"`
			M4Afilesize   int    `json:"m4afilesize"`
			RpPublish     int    `json:"rp_publish"`
			Privilege     int    `json:"privilege"`
			Filesize1     int    `json:"filesize"`
			Isoriginal    int    `json:"isoriginal"`
			Privilege1    int    `json:"320privilege"`
			Sqprivilege   int    `json:"sqprivilege"`
			FailProcessSq int    `json:"fail_process_sq"`
		} `json:"group"`
		Price320       int    `json:"price_320"`
		Hash1          string `json:"320hash"`
		Topic          string `json:"topic"`
		Othername      string `json:"othername"`
		Isnew          int    `json:"isnew"`
		FoldType       int    `json:"fold_type"`
		OldCpy         int    `json:"old_cpy"`
		Srctype        int    `json:"srctype"`
		Singername     string `json:"singername"`
		AlbumAudioId   int    `json:"album_audio_id"`
		Duration       int    `json:"duration"`
		Filesize       int    `json:"320filesize"`
		PkgPrice320    int    `json:"pkg_price_320"`
		AudioId        int    `json:"audio_id"`
		Feetype        int    `json:"feetype"`
		Price          int    `json:"price"`
		Filename       string `json:"filename"`
		Source         string `json:"source"`
		PriceSq        int    `json:"price_sq"`
		FailProcess320 int    `json:"fail_process_320"`
		TransParam     struct {
			CpyLevel int `json:"cpy_level"`
			CpyGrade int `json:"cpy_grade"`
			Classmap struct {
				Attr0 int64 `json:"attr0"`
			} `json:"classmap"`
			PayBlockTpl int `json:"pay_block_tpl"`
			Qualitymap  struct {
				Attr0 int `json:"attr0"`
			} `json:"qualitymap"`
			AppidBlock     string `json:"appid_block,omitempty"`
			HashMultitrack string `json:"hash_multitrack"`
			Language       string `json:"language"`
			Cid            int    `json:"cid"`
			CpyAttr0       int    `json:"cpy_attr0"`
			Ipmap          struct {
				Attr0 int `json:"attr0"`
			} `json:"ipmap,omitempty"`
			HashOffset struct {
				ClipHash   string `json:"clip_hash"`
				StartByte  int    `json:"start_byte"`
				EndMs      int    `json:"end_ms"`
				EndByte    int    `json:"end_byte"`
				FileType   int    `json:"file_type"`
				StartMs    int    `json:"start_ms"`
				OffsetHash string `json:"offset_hash"`
			} `json:"hash_offset,omitempty"`
			MusicpackAdvance int    `json:"musicpack_advance"`
			Display          int    `json:"display"`
			DisplayRate      int    `json:"display_rate"`
			FreeForAd        int    `json:"free_for_ad,omitempty"`
			SongnameSuffix   string `json:"songname_suffix,omitempty"`
		} `json:"trans_param"`
		PkgPrice      int    `json:"pkg_price"`
		PayType320    int    `json:"pay_type_320"`
		TopicUrl      string `json:"topic_url"`
		M4Afilesize   int    `json:"m4afilesize"`
		RpPublish     int    `json:"rp_publish"`
		Privilege     int    `json:"privilege"`
		Filesize1     int    `json:"filesize"`
		Isoriginal    int    `json:"isoriginal"`
		Privilege1    int    `json:"320privilege"`
		Sqprivilege   int    `json:"sqprivilege"`
		FailProcessSq int    `json:"fail_process_sq"`
	} `json:"info"`
	Aggregation   []interface{} `json:"aggregation"`
	Correctiontip string        `json:"correctiontip"`
	Istagresult   int           `json:"istagresult"`
}

type KugouSearchResp struct {
	Status  int             `json:"status"`
	Errcode int             `json:"errcode"`
	Data    KugouSearchData `json:"data"`
	Error   string          `json:"error"`
}

type KugouInfoResp struct {
	FileHead int `json:"fileHead"`
	Q        int `json:"q"`
	Extra    struct {
		Filesize       int    `json:"320filesize"`
		Sqfilesize     int    `json:"sqfilesize"`
		Sqhash         string `json:"sqhash"`
		Sqtimelength   int    `json:"sqtimelength"`
		Hash           string `json:"320hash"`
		Filesize1      int    `json:"128filesize"`
		Timelength     int    `json:"128timelength"`
		Timelength1    int    `json:"320timelength"`
		Sqbitrate      int    `json:"sqbitrate"`
		Highbitrate    int    `json:"highbitrate"`
		Highfilesize   int    `json:"highfilesize"`
		Hash1          string `json:"128hash"`
		Bitrate        int    `json:"128bitrate"`
		Hightimelength int    `json:"hightimelength"`
		Highhash       string `json:"highhash"`
	} `json:"extra"`
	FileSize     int    `json:"fileSize"`
	ChoricSinger string `json:"choricSinger"`
	Authors      []struct {
		Identity   int    `json:"identity"`
		Avatar     string `json:"avatar"`
		AuthorId   int    `json:"author_id"`
		Birthday   string `json:"birthday"`
		Language   string `json:"language"`
		Type       int    `json:"type"`
		IsPublish  int    `json:"is_publish"`
		Country    string `json:"country"`
		AuthorName string `json:"author_name"`
	} `json:"authors"`
	AlbumImg   string `json:"album_img"`
	ReqAlbumid string `json:"req_albumid"`
	Url        string `json:"url"`
	Time       int    `json:"time"`
	TransParam struct {
		CpyGrade         int `json:"cpy_grade"`
		MusicpackAdvance int `json:"musicpack_advance"`
		Qualitymap       struct {
			Attr0 int `json:"attr0"`
		} `json:"qualitymap"`
		CpyLevel int `json:"cpy_level"`
		CpyAttr0 int `json:"cpy_attr0"`
		Ipmap    struct {
			Attr0 int `json:"attr0"`
		} `json:"ipmap"`
		Cid            int    `json:"cid"`
		HashMultitrack string `json:"hash_multitrack"`
		PayBlockTpl    int    `json:"pay_block_tpl"`
		Classmap       struct {
			Attr0 int `json:"attr0"`
		} `json:"classmap"`
		DisplayRate int    `json:"display_rate"`
		AppidBlock  string `json:"appid_block"`
		Display     int    `json:"display"`
		Language    string `json:"language"`
	} `json:"trans_param"`
	OldCpy          int         `json:"old_cpy"`
	Albumid         int         `json:"albumid"`
	SingerName      string      `json:"singerName"`
	TopicUrl        string      `json:"topic_url"`
	ExtName         string      `json:"extName"`
	AuthorName      string      `json:"author_name"`
	Privilege       int         `json:"privilege"`
	PayType         int         `json:"pay_type"`
	SongName        string      `json:"songName"`
	SingerHead      string      `json:"singerHead"`
	Mvhash          string      `json:"mvhash"`
	Hash            string      `json:"hash"`
	AudioId         int         `json:"audio_id"`
	StoreType       string      `json:"store_type"`
	Error           string      `json:"error"`
	ImgUrl          string      `json:"imgUrl"`
	Highprivilege   int         `json:"highprivilege"`
	AlbumAudioId    int         `json:"album_audio_id"`
	AreaCode        string      `json:"area_code"`
	Errcode         int         `json:"errcode"`
	BitRate         int         `json:"bitRate"`
	Privilege1      int         `json:"128privilege"`
	SingerId        int         `json:"singerId"`
	AudioGroupId    int         `json:"audio_group_id"`
	ReqHash         string      `json:"req_hash"`
	PublishSelfFlag int         `json:"publish_self_flag"`
	Status          int         `json:"status"`
	Stype           int         `json:"stype"`
	Privilege2      int         `json:"320privilege"`
	FailProcess     int         `json:"fail_process"`
	Ctype           int         `json:"ctype"`
	FileName        string      `json:"fileName"`
	TopicRemark     string      `json:"topic_remark"`
	Intro           string      `json:"intro"`
	AlbumCategory   int         `json:"album_category"`
	Sqprivilege     int         `json:"sqprivilege"`
	BackupUrl       interface{} `json:"backup_url"`
	TimeLength      int         `json:"timeLength"`
}
