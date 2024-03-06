package models

import (
	"fmt"
	"github.com/spf13/cast"
	"net/url"
	"path/filepath"
	"time"
)

// VideoViewResp
// https://socialsisteryi.github.io/bilibili-API-collect/docs/video/info.html#%E8%8E%B7%E5%8F%96%E8%A7%86%E9%A2%91%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF-web%E7%AB%AF
// 根对象：
// 字段	类型	内容	备注
// code	num	返回值	0：成功
// -400：请求错误
// -403：权限不足
// -404：无视频
// 62002：稿件不可见
// 62004：稿件审核中
// message	str	错误信息	默认为0
// ttl	num	1
// data	obj	信息本体
type VideoViewResp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Ttl     int           `json:"ttl"`
	Data    VideoViewData `json:"data"`
}

type VideoViewData struct {
	Bvid      string `json:"bvid"`
	Aid       int    `json:"aid"`
	Videos    int    `json:"videos"`
	Tid       int    `json:"tid"`
	Tname     string `json:"tname"`
	Copyright int    `json:"copyright"`
	Pic       string `json:"pic"`
	Title     string `json:"title"`
	Pubdate   int    `json:"pubdate"`
	Ctime     int    `json:"ctime"`
	Desc      string `json:"desc"`
	DescV2    []struct {
		RawText string `json:"raw_text"`
		Type    int    `json:"type"`
		BizId   int    `json:"biz_id"`
	} `json:"desc_v2"`
	State    int `json:"state"`
	Duration int `json:"duration"`
	Rights   struct {
		Bp            int `json:"bp"`
		Elec          int `json:"elec"`
		Download      int `json:"download"`
		Movie         int `json:"movie"`
		Pay           int `json:"pay"`
		Hd5           int `json:"hd5"`
		NoReprint     int `json:"no_reprint"`
		Autoplay      int `json:"autoplay"`
		UgcPay        int `json:"ugc_pay"`
		IsCooperation int `json:"is_cooperation"`
		UgcPayPreview int `json:"ugc_pay_preview"`
		NoBackground  int `json:"no_background"`
		CleanMode     int `json:"clean_mode"`
		IsSteinGate   int `json:"is_stein_gate"`
		Is360         int `json:"is_360"`
		NoShare       int `json:"no_share"`
		ArcPay        int `json:"arc_pay"`
		FreeWatch     int `json:"free_watch"`
	} `json:"rights"`
	Owner struct {
		Mid  int    `json:"mid"`
		Name string `json:"name"`
		Face string `json:"face"`
	} `json:"owner"`
	Stat struct {
		Aid        int    `json:"aid"`
		View       int    `json:"view"`
		Danmaku    int    `json:"danmaku"`
		Reply      int    `json:"reply"`
		Favorite   int    `json:"favorite"`
		Coin       int    `json:"coin"`
		Share      int    `json:"share"`
		NowRank    int    `json:"now_rank"`
		HisRank    int    `json:"his_rank"`
		Like       int    `json:"like"`
		Dislike    int    `json:"dislike"`
		Evaluation string `json:"evaluation"`
		Vt         int    `json:"vt"`
	} `json:"stat"`
	ArgueInfo struct {
		ArgueMsg  string `json:"argue_msg"`
		ArgueType int    `json:"argue_type"`
		ArgueLink string `json:"argue_link"`
	} `json:"argue_info"`
	Dynamic   string `json:"dynamic"`
	Cid       int    `json:"cid"`
	Dimension struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		Rotate int `json:"rotate"`
	} `json:"dimension"`
	SeasonId           int         `json:"season_id"`
	Premiere           interface{} `json:"premiere"`
	TeenageMode        int         `json:"teenage_mode"`
	IsChargeableSeason bool        `json:"is_chargeable_season"`
	IsStory            bool        `json:"is_story"`
	IsUpowerExclusive  bool        `json:"is_upower_exclusive"`
	IsUpowerPlay       bool        `json:"is_upower_play"`
	IsUpowerPreview    bool        `json:"is_upower_preview"`
	EnableVt           int         `json:"enable_vt"`
	VtDisplay          string      `json:"vt_display"`
	NoCache            bool        `json:"no_cache"`
	Pages              []struct {
		Cid       int    `json:"cid"`
		Page      int    `json:"page"`
		From      string `json:"from"`
		Part      string `json:"part"`
		Duration  int    `json:"duration"`
		Vid       string `json:"vid"`
		Weblink   string `json:"weblink"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		FirstFrame string `json:"first_frame"`
	} `json:"pages"`
	Subtitle struct {
		AllowSubmit bool          `json:"allow_submit"`
		List        []interface{} `json:"list"`
	} `json:"subtitle"`
	UgcSeason struct {
		Id        int    `json:"id"`
		Title     string `json:"title"`
		Cover     string `json:"cover"`
		Mid       int    `json:"mid"`
		Intro     string `json:"intro"`
		SignState int    `json:"sign_state"`
		Attribute int    `json:"attribute"`
		Sections  []struct {
			SeasonId int    `json:"season_id"`
			Id       int    `json:"id"`
			Title    string `json:"title"`
			Type     int    `json:"type"`
			Episodes []struct {
				SeasonId  int    `json:"season_id"`
				SectionId int    `json:"section_id"`
				Id        int    `json:"id"`
				Aid       int    `json:"aid"`
				Cid       int    `json:"cid"`
				Title     string `json:"title"`
				Attribute int    `json:"attribute"`
				Arc       struct {
					Aid       int    `json:"aid"`
					Videos    int    `json:"videos"`
					TypeId    int    `json:"type_id"`
					TypeName  string `json:"type_name"`
					Copyright int    `json:"copyright"`
					Pic       string `json:"pic"`
					Title     string `json:"title"`
					Pubdate   int    `json:"pubdate"`
					Ctime     int    `json:"ctime"`
					Desc      string `json:"desc"`
					State     int    `json:"state"`
					Duration  int    `json:"duration"`
					Rights    struct {
						Bp            int `json:"bp"`
						Elec          int `json:"elec"`
						Download      int `json:"download"`
						Movie         int `json:"movie"`
						Pay           int `json:"pay"`
						Hd5           int `json:"hd5"`
						NoReprint     int `json:"no_reprint"`
						Autoplay      int `json:"autoplay"`
						UgcPay        int `json:"ugc_pay"`
						IsCooperation int `json:"is_cooperation"`
						UgcPayPreview int `json:"ugc_pay_preview"`
						ArcPay        int `json:"arc_pay"`
						FreeWatch     int `json:"free_watch"`
					} `json:"rights"`
					Author struct {
						Mid  int    `json:"mid"`
						Name string `json:"name"`
						Face string `json:"face"`
					} `json:"author"`
					Stat struct {
						Aid        int    `json:"aid"`
						View       int    `json:"view"`
						Danmaku    int    `json:"danmaku"`
						Reply      int    `json:"reply"`
						Fav        int    `json:"fav"`
						Coin       int    `json:"coin"`
						Share      int    `json:"share"`
						NowRank    int    `json:"now_rank"`
						HisRank    int    `json:"his_rank"`
						Like       int    `json:"like"`
						Dislike    int    `json:"dislike"`
						Evaluation string `json:"evaluation"`
						ArgueMsg   string `json:"argue_msg"`
						Vt         int    `json:"vt"`
						Vv         int    `json:"vv"`
					} `json:"stat"`
					Dynamic   string `json:"dynamic"`
					Dimension struct {
						Width  int `json:"width"`
						Height int `json:"height"`
						Rotate int `json:"rotate"`
					} `json:"dimension"`
					DescV2             interface{} `json:"desc_v2"`
					IsChargeableSeason bool        `json:"is_chargeable_season"`
					IsBlooper          bool        `json:"is_blooper"`
					EnableVt           int         `json:"enable_vt"`
					VtDisplay          string      `json:"vt_display"`
				} `json:"arc"`
				Page struct {
					Cid       int    `json:"cid"`
					Page      int    `json:"page"`
					From      string `json:"from"`
					Part      string `json:"part"`
					Duration  int    `json:"duration"`
					Vid       string `json:"vid"`
					Weblink   string `json:"weblink"`
					Dimension struct {
						Width  int `json:"width"`
						Height int `json:"height"`
						Rotate int `json:"rotate"`
					} `json:"dimension"`
				} `json:"page"`
				Bvid string `json:"bvid"`
			} `json:"episodes"`
		} `json:"sections"`
		Stat struct {
			SeasonId int `json:"season_id"`
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Fav      int `json:"fav"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Vt       int `json:"vt"`
			Vv       int `json:"vv"`
		} `json:"stat"`
		EpCount     int  `json:"ep_count"`
		SeasonType  int  `json:"season_type"`
		IsPaySeason bool `json:"is_pay_season"`
		EnableVt    int  `json:"enable_vt"`
	} `json:"ugc_season"`
	IsSeasonDisplay bool `json:"is_season_display"`
	UserGarb        struct {
		UrlImageAniCut string `json:"url_image_ani_cut"`
	} `json:"user_garb"`
	HonorReply struct {
		Honor []struct {
			Aid                int    `json:"aid"`
			Type               int    `json:"type"`
			Desc               string `json:"desc"`
			WeeklyRecommendNum int    `json:"weekly_recommend_num"`
		} `json:"honor"`
	} `json:"honor_reply"`
	LikeIcon          string `json:"like_icon"`
	NeedJumpBv        bool   `json:"need_jump_bv"`
	DisableShowUpInfo bool   `json:"disable_show_up_info"`
	IsStoryPlay       int    `json:"is_story_play"`
}

// VideoViewDetailResp
// https://socialsisteryi.github.io/bilibili-API-collect/docs/video/info.html#%E8%8E%B7%E5%8F%96%E8%A7%86%E9%A2%91%E8%B6%85%E8%AF%A6%E7%BB%86%E4%BF%A1%E6%81%AF-web%E7%AB%AF
// 根对象：
// 字段	类型	内容	备注
// code	num	返回值	0：成功
// -400：请求错误
// -403：权限不足
// -404：无视频
// 62002：稿件不可见
// 62004：稿件审核中
// message	str	错误信息	默认为0
// ttl	num	1
// data	obj	信息本体
type VideoViewDetailResp struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Ttl     int                 `json:"ttl"`
	Data    VideoViewDetailData `json:"data"`
}

type VideoViewDetailData struct {
	View struct {
		Bvid      string `json:"bvid"`
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		DescV2    []struct {
			RawText string `json:"raw_text"`
			Type    int    `json:"type"`
			BizId   int    `json:"biz_id"`
		} `json:"desc_v2"`
		State    int `json:"state"`
		Duration int `json:"duration"`
		Rights   struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			CleanMode     int `json:"clean_mode"`
			IsSteinGate   int `json:"is_stein_gate"`
			Is360         int `json:"is_360"`
			NoShare       int `json:"no_share"`
			ArcPay        int `json:"arc_pay"`
			FreeWatch     int `json:"free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid        int    `json:"aid"`
			View       int    `json:"view"`
			Danmaku    int    `json:"danmaku"`
			Reply      int    `json:"reply"`
			Favorite   int    `json:"favorite"`
			Coin       int    `json:"coin"`
			Share      int    `json:"share"`
			NowRank    int    `json:"now_rank"`
			HisRank    int    `json:"his_rank"`
			Like       int    `json:"like"`
			Dislike    int    `json:"dislike"`
			Evaluation string `json:"evaluation"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		ArgueInfo struct {
			ArgueMsg  string `json:"argue_msg"`
			ArgueType int    `json:"argue_type"`
			ArgueLink string `json:"argue_link"`
		} `json:"argue_info"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonId           int         `json:"season_id"`
		Premiere           interface{} `json:"premiere"`
		TeenageMode        int         `json:"teenage_mode"`
		IsChargeableSeason bool        `json:"is_chargeable_season"`
		IsStory            bool        `json:"is_story"`
		IsUpowerExclusive  bool        `json:"is_upower_exclusive"`
		IsUpowerPlay       bool        `json:"is_upower_play"`
		IsUpowerPreview    bool        `json:"is_upower_preview"`
		EnableVt           int         `json:"enable_vt"`
		VtDisplay          string      `json:"vt_display"`
		NoCache            bool        `json:"no_cache"`
		Pages              []struct {
			Cid       int    `json:"cid"`
			Page      int    `json:"page"`
			From      string `json:"from"`
			Part      string `json:"part"`
			Duration  int    `json:"duration"`
			Vid       string `json:"vid"`
			Weblink   string `json:"weblink"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			FirstFrame string `json:"first_frame"`
		} `json:"pages"`
		Subtitle struct {
			AllowSubmit bool          `json:"allow_submit"`
			List        []interface{} `json:"list"`
		} `json:"subtitle"`
		UgcSeason struct {
			Id        int    `json:"id"`
			Title     string `json:"title"`
			Cover     string `json:"cover"`
			Mid       int    `json:"mid"`
			Intro     string `json:"intro"`
			SignState int    `json:"sign_state"`
			Attribute int    `json:"attribute"`
			Sections  []struct {
				SeasonId int    `json:"season_id"`
				Id       int    `json:"id"`
				Title    string `json:"title"`
				Type     int    `json:"type"`
				Episodes []struct {
					SeasonId  int    `json:"season_id"`
					SectionId int    `json:"section_id"`
					Id        int    `json:"id"`
					Aid       int    `json:"aid"`
					Cid       int    `json:"cid"`
					Title     string `json:"title"`
					Attribute int    `json:"attribute"`
					Arc       struct {
						Aid       int    `json:"aid"`
						Videos    int    `json:"videos"`
						TypeId    int    `json:"type_id"`
						TypeName  string `json:"type_name"`
						Copyright int    `json:"copyright"`
						Pic       string `json:"pic"`
						Title     string `json:"title"`
						Pubdate   int    `json:"pubdate"`
						Ctime     int    `json:"ctime"`
						Desc      string `json:"desc"`
						State     int    `json:"state"`
						Duration  int    `json:"duration"`
						Rights    struct {
							Bp            int `json:"bp"`
							Elec          int `json:"elec"`
							Download      int `json:"download"`
							Movie         int `json:"movie"`
							Pay           int `json:"pay"`
							Hd5           int `json:"hd5"`
							NoReprint     int `json:"no_reprint"`
							Autoplay      int `json:"autoplay"`
							UgcPay        int `json:"ugc_pay"`
							IsCooperation int `json:"is_cooperation"`
							UgcPayPreview int `json:"ugc_pay_preview"`
							ArcPay        int `json:"arc_pay"`
							FreeWatch     int `json:"free_watch"`
						} `json:"rights"`
						Author struct {
							Mid  int    `json:"mid"`
							Name string `json:"name"`
							Face string `json:"face"`
						} `json:"author"`
						Stat struct {
							Aid        int    `json:"aid"`
							View       int    `json:"view"`
							Danmaku    int    `json:"danmaku"`
							Reply      int    `json:"reply"`
							Fav        int    `json:"fav"`
							Coin       int    `json:"coin"`
							Share      int    `json:"share"`
							NowRank    int    `json:"now_rank"`
							HisRank    int    `json:"his_rank"`
							Like       int    `json:"like"`
							Dislike    int    `json:"dislike"`
							Evaluation string `json:"evaluation"`
							ArgueMsg   string `json:"argue_msg"`
							Vt         int    `json:"vt"`
							Vv         int    `json:"vv"`
						} `json:"stat"`
						Dynamic   string `json:"dynamic"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
						DescV2             interface{} `json:"desc_v2"`
						IsChargeableSeason bool        `json:"is_chargeable_season"`
						IsBlooper          bool        `json:"is_blooper"`
						EnableVt           int         `json:"enable_vt"`
						VtDisplay          string      `json:"vt_display"`
					} `json:"arc"`
					Page struct {
						Cid       int    `json:"cid"`
						Page      int    `json:"page"`
						From      string `json:"from"`
						Part      string `json:"part"`
						Duration  int    `json:"duration"`
						Vid       string `json:"vid"`
						Weblink   string `json:"weblink"`
						Dimension struct {
							Width  int `json:"width"`
							Height int `json:"height"`
							Rotate int `json:"rotate"`
						} `json:"dimension"`
					} `json:"page"`
					Bvid string `json:"bvid"`
				} `json:"episodes"`
			} `json:"sections"`
			Stat struct {
				SeasonId int `json:"season_id"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Fav      int `json:"fav"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			EpCount     int  `json:"ep_count"`
			SeasonType  int  `json:"season_type"`
			IsPaySeason bool `json:"is_pay_season"`
			EnableVt    int  `json:"enable_vt"`
		} `json:"ugc_season"`
		IsSeasonDisplay bool `json:"is_season_display"`
		UserGarb        struct {
			UrlImageAniCut string `json:"url_image_ani_cut"`
		} `json:"user_garb"`
		HonorReply struct {
			Honor []struct {
				Aid                int    `json:"aid"`
				Type               int    `json:"type"`
				Desc               string `json:"desc"`
				WeeklyRecommendNum int    `json:"weekly_recommend_num"`
			} `json:"honor"`
		} `json:"honor_reply"`
		LikeIcon          string `json:"like_icon"`
		NeedJumpBv        bool   `json:"need_jump_bv"`
		DisableShowUpInfo bool   `json:"disable_show_up_info"`
		IsStoryPlay       int    `json:"is_story_play"`
	} `json:"View"`
	Card struct {
		Card struct {
			Mid         string        `json:"mid"`
			Name        string        `json:"name"`
			Approve     bool          `json:"approve"`
			Sex         string        `json:"sex"`
			Rank        string        `json:"rank"`
			Face        string        `json:"face"`
			FaceNft     int           `json:"face_nft"`
			FaceNftType int           `json:"face_nft_type"`
			DisplayRank string        `json:"DisplayRank"`
			Regtime     int           `json:"regtime"`
			Spacesta    int           `json:"spacesta"`
			Birthday    string        `json:"birthday"`
			Place       string        `json:"place"`
			Description string        `json:"description"`
			Article     int           `json:"article"`
			Attentions  []interface{} `json:"attentions"`
			Fans        int           `json:"fans"`
			Friend      int           `json:"friend"`
			Attention   int           `json:"attention"`
			Sign        string        `json:"sign"`
			LevelInfo   struct {
				CurrentLevel int `json:"current_level"`
				CurrentMin   int `json:"current_min"`
				CurrentExp   int `json:"current_exp"`
				NextExp      int `json:"next_exp"`
			} `json:"level_info"`
			Pendant struct {
				Pid               int    `json:"pid"`
				Name              string `json:"name"`
				Image             string `json:"image"`
				Expire            int    `json:"expire"`
				ImageEnhance      string `json:"image_enhance"`
				ImageEnhanceFrame string `json:"image_enhance_frame"`
				NPid              int    `json:"n_pid"`
			} `json:"pendant"`
			Nameplate struct {
				Nid        int    `json:"nid"`
				Name       string `json:"name"`
				Image      string `json:"image"`
				ImageSmall string `json:"image_small"`
				Level      string `json:"level"`
				Condition  string `json:"condition"`
			} `json:"nameplate"`
			Official struct {
				Role  int    `json:"role"`
				Title string `json:"title"`
				Desc  string `json:"desc"`
				Type  int    `json:"type"`
			} `json:"Official"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify"`
			Vip struct {
				Type       int   `json:"type"`
				Status     int   `json:"status"`
				DueDate    int64 `json:"due_date"`
				VipPayType int   `json:"vip_pay_type"`
				ThemeType  int   `json:"theme_type"`
				Label      struct {
					Path                  string `json:"path"`
					Text                  string `json:"text"`
					LabelTheme            string `json:"label_theme"`
					TextColor             string `json:"text_color"`
					BgStyle               int    `json:"bg_style"`
					BgColor               string `json:"bg_color"`
					BorderColor           string `json:"border_color"`
					UseImgLabel           bool   `json:"use_img_label"`
					ImgLabelUriHans       string `json:"img_label_uri_hans"`
					ImgLabelUriHant       string `json:"img_label_uri_hant"`
					ImgLabelUriHansStatic string `json:"img_label_uri_hans_static"`
					ImgLabelUriHantStatic string `json:"img_label_uri_hant_static"`
				} `json:"label"`
				AvatarSubscript    int    `json:"avatar_subscript"`
				NicknameColor      string `json:"nickname_color"`
				Role               int    `json:"role"`
				AvatarSubscriptUrl string `json:"avatar_subscript_url"`
				TvVipStatus        int    `json:"tv_vip_status"`
				TvVipPayType       int    `json:"tv_vip_pay_type"`
				TvDueDate          int    `json:"tv_due_date"`
				AvatarIcon         struct {
					IconType     int `json:"icon_type"`
					IconResource struct {
					} `json:"icon_resource"`
				} `json:"avatar_icon"`
				VipType   int `json:"vipType"`
				VipStatus int `json:"vipStatus"`
			} `json:"vip"`
			IsSeniorMember int `json:"is_senior_member"`
		} `json:"card"`
		Space struct {
			SImg string `json:"s_img"`
			LImg string `json:"l_img"`
		} `json:"space"`
		Following    bool `json:"following"`
		ArchiveCount int  `json:"archive_count"`
		ArticleCount int  `json:"article_count"`
		Follower     int  `json:"follower"`
		LikeNum      int  `json:"like_num"`
	} `json:"Card"`
	Tags []struct {
		TagId        int    `json:"tag_id"`
		TagName      string `json:"tag_name"`
		Cover        string `json:"cover"`
		HeadCover    string `json:"head_cover"`
		Content      string `json:"content"`
		ShortContent string `json:"short_content"`
		Type         int    `json:"type"`
		State        int    `json:"state"`
		Ctime        int    `json:"ctime"`
		Count        struct {
			View  int `json:"view"`
			Use   int `json:"use"`
			Atten int `json:"atten"`
		} `json:"count"`
		IsAtten         int    `json:"is_atten"`
		Likes           int    `json:"likes"`
		Hates           int    `json:"hates"`
		Attribute       int    `json:"attribute"`
		Liked           int    `json:"liked"`
		Hated           int    `json:"hated"`
		ExtraAttr       int    `json:"extra_attr"`
		MusicId         string `json:"music_id"`
		TagType         string `json:"tag_type"`
		IsActivity      bool   `json:"is_activity"`
		Color           string `json:"color"`
		Alpha           int    `json:"alpha"`
		IsSeason        bool   `json:"is_season"`
		SubscribedCount int    `json:"subscribed_count"`
		ArchiveCount    string `json:"archive_count"`
		FeaturedCount   int    `json:"featured_count"`
		JumpUrl         string `json:"jump_url"`
	} `json:"Tags"`
	Reply struct {
		Page    interface{} `json:"page"`
		Replies []struct {
			Rpid       int         `json:"rpid"`
			Oid        int         `json:"oid"`
			Type       int         `json:"type"`
			Mid        int         `json:"mid"`
			Root       int         `json:"root"`
			Parent     int         `json:"parent"`
			Dialog     int         `json:"dialog"`
			Count      int         `json:"count"`
			Rcount     int         `json:"rcount"`
			State      int         `json:"state"`
			Fansgrade  int         `json:"fansgrade"`
			Attr       int         `json:"attr"`
			Ctime      int         `json:"ctime"`
			Like       int         `json:"like"`
			Action     int         `json:"action"`
			Content    interface{} `json:"content"`
			Replies    interface{} `json:"replies"`
			Assist     int         `json:"assist"`
			ShowFollow bool        `json:"show_follow"`
		} `json:"replies"`
	} `json:"Reply"`
	Related []struct {
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			ArcPay        int `json:"arc_pay"`
			PayFreeWatch  int `json:"pay_free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid      int `json:"aid"`
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Favorite int `json:"favorite"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Dislike  int `json:"dislike"`
			Vt       int `json:"vt"`
			Vv       int `json:"vv"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		SeasonId    int         `json:"season_id,omitempty"`
		ShortLinkV2 string      `json:"short_link_v2"`
		FirstFrame  string      `json:"first_frame,omitempty"`
		Bvid        string      `json:"bvid"`
		SeasonType  int         `json:"season_type"`
		IsOgv       bool        `json:"is_ogv"`
		OgvInfo     interface{} `json:"ogv_info"`
		RcmdReason  string      `json:"rcmd_reason"`
		EnableVt    int         `json:"enable_vt"`
		AiRcmd      struct {
			Id      int    `json:"id"`
			Goto    string `json:"goto"`
			Trackid string `json:"trackid"`
			UniqId  string `json:"uniq_id"`
		} `json:"ai_rcmd"`
		PubLocation string `json:"pub_location,omitempty"`
		UpFromV2    int    `json:"up_from_v2,omitempty"`
		MissionId   int    `json:"mission_id,omitempty"`
		RedirectUrl string `json:"redirect_url,omitempty"`
	} `json:"Related"`
	Spec     interface{} `json:"Spec"`
	HotShare struct {
		Show bool          `json:"show"`
		List []interface{} `json:"list"`
	} `json:"hot_share"`
	Elec      interface{} `json:"elec"`
	Recommend interface{} `json:"recommend"`
	Emergency struct {
		NoLike  bool `json:"no_like"`
		NoCoin  bool `json:"no_coin"`
		NoFav   bool `json:"no_fav"`
		NoShare bool `json:"no_share"`
	} `json:"emergency"`
	ViewAddit struct {
		Field1 bool `json:"63"`
		Field2 bool `json:"64"`
		Field3 bool `json:"69"`
		Field4 bool `json:"71"`
		Field5 bool `json:"72"`
	} `json:"view_addit"`
	Guide      interface{} `json:"guide"`
	QueryTags  interface{} `json:"query_tags"`
	IsOldUser  bool        `json:"is_old_user"`
	Participle []string    `json:"participle"`
}

type VideoPlayerWbiPlayUrlParams struct {
	Avid int    //	稿件 avid	必要（可选）	avid 与 bvid 任选一个
	Bvid string //	稿件 bvid	必要（可选）	avid 与 bvid 任选一个
	Cid  int    //	视频 cid	必要
	//含义见 上表
	//DASH 格式时无效
	Qn int //	视频清晰度选择	非必要	未登录默认 32（480P），登录后默认 64（720P）
	//含义见 上表
	Fnval int //	视频流格式标识	非必要	默认值为1（MP4 格式）
	Fnver int //	0	非必要
	//画质最高 4K：1
	Fourk   int    //	是否允许 4K 视频	非必要	画质最高 1080P：0（默认）
	Session string //		非必要	从视频播放页的 HTML 中获取
	OType   string //		非必要	固定为json
	Type    string //		非必要	目前为空
	//html5：移动端 HTML5 播放（仅支持 MP4 格式，无 referer 鉴权可以直接使用video标签播放）
	Platform    string //		非必要	pc：web播放（默认值，视频流存在 referer鉴权）
	HighQuality int    //	是否高画质	非必要	platform=html5时，此值为1可使画质为1080p
}

// VideoPlayerWbiPlayUrlResp
// https://socialsisteryi.github.io/bilibili-API-collect/docs/video/videostream_url.html#%E8%8E%B7%E5%8F%96%E8%A7%86%E9%A2%91%E6%B5%81%E5%9C%B0%E5%9D%80-web%E7%AB%AF
// 根对象：
// 字段	类型	内容	备注
// code	num	返回值	0：成功
// -400：请求错误
// -404：无视频
// message	str	错误信息	默认为0
// ttl	num	1
// data	有效时：obj
// 无效时：null	数据本体
type VideoPlayerWbiPlayUrlResp struct {
	Code    int                       `json:"code"`
	Message string                    `json:"message"`
	Ttl     int                       `json:"ttl"`
	Data    VideoPlayerWbiPlayUrlData `json:"data"`
}

type VideoPlayerWbiPlayUrlData struct {
	From              string   `json:"from"`
	Result            string   `json:"result"`
	Message           string   `json:"message"`
	Quality           int      `json:"quality"`
	Format            string   `json:"format"`
	Timelength        int      `json:"timelength"`
	AcceptFormat      string   `json:"accept_format"`
	AcceptDescription []string `json:"accept_description"`
	AcceptQuality     []int    `json:"accept_quality"`
	VideoCodecid      int      `json:"video_codecid"`
	SeekParam         string   `json:"seek_param"`
	SeekType          string   `json:"seek_type"`
	Durl              []struct {
		Order     int      `json:"order"`
		Length    int      `json:"length"`
		Size      int      `json:"size"`
		Ahead     string   `json:"ahead"`
		Vhead     string   `json:"vhead"`
		Url       string   `json:"url"`
		BackupUrl []string `json:"backup_url"`
	} `json:"durl"`
	Dash struct {
		Duration       int     `json:"duration"`
		MinBufferTime  float64 `json:"minBufferTime"`
		MinBufferTime1 float64 `json:"min_buffer_time"`
		Video          []struct {
			Id            int      `json:"id"`
			BaseUrl       string   `json:"baseUrl"`
			BaseUrl1      string   `json:"base_url"`
			BackupUrl     []string `json:"backupUrl"`
			BackupUrl1    []string `json:"backup_url"`
			Bandwidth     int      `json:"bandwidth"`
			MimeType      string   `json:"mimeType"`
			MimeType1     string   `json:"mime_type"`
			Codecs        string   `json:"codecs"`
			Width         int      `json:"width"`
			Height        int      `json:"height"`
			FrameRate     string   `json:"frameRate"`
			FrameRate1    string   `json:"frame_rate"`
			Sar           string   `json:"sar"`
			StartWithSap  int      `json:"startWithSap"`
			StartWithSap1 int      `json:"start_with_sap"`
			SegmentBase   struct {
				Initialization string `json:"Initialization"`
				IndexRange     string `json:"indexRange"`
			} `json:"SegmentBase"`
			SegmentBase1 struct {
				Initialization string `json:"initialization"`
				IndexRange     string `json:"index_range"`
			} `json:"segment_base"`
			Codecid int `json:"codecid"`
		} `json:"video"`
		Audio []struct {
			Id            int      `json:"id"`
			BaseUrl       string   `json:"baseUrl"`
			BaseUrl1      string   `json:"base_url"`
			BackupUrl     []string `json:"backupUrl"`
			BackupUrl1    []string `json:"backup_url"`
			Bandwidth     int      `json:"bandwidth"`
			MimeType      string   `json:"mimeType"`
			MimeType1     string   `json:"mime_type"`
			Codecs        string   `json:"codecs"`
			Width         int      `json:"width"`
			Height        int      `json:"height"`
			FrameRate     string   `json:"frameRate"`
			FrameRate1    string   `json:"frame_rate"`
			Sar           string   `json:"sar"`
			StartWithSap  int      `json:"startWithSap"`
			StartWithSap1 int      `json:"start_with_sap"`
			SegmentBase   struct {
				Initialization string `json:"Initialization"`
				IndexRange     string `json:"indexRange"`
			} `json:"SegmentBase"`
			SegmentBase1 struct {
				Initialization string `json:"initialization"`
				IndexRange     string `json:"index_range"`
			} `json:"segment_base"`
			Codecid int `json:"codecid"`
		} `json:"audio"`
		Dolby struct {
			Type  int         `json:"type"`
			Audio interface{} `json:"audio"`
		} `json:"dolby"`
		Flac interface{} `json:"flac"`
	} `json:"dash"`
	SupportFormats []struct {
		Quality        int      `json:"quality"`
		Format         string   `json:"format"`
		NewDescription string   `json:"new_description"`
		DisplayDesc    string   `json:"display_desc"`
		Superscript    string   `json:"superscript"`
		Codecs         []string `json:"codecs"`
	} `json:"support_formats"`
	HighFormat   interface{} `json:"high_format"`
	LastPlayTime int         `json:"last_play_time"`
	LastPlayCid  int         `json:"last_play_cid"`
}

type VideoRankingResp struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Ttl     int              `json:"ttl"`
	Data    VideoRankingData `json:"data"`
}

type VideoRankingData struct {
	Note string `json:"note"`
	List []struct {
		Aid       int    `json:"aid"`
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"`
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		MissionId int    `json:"mission_id,omitempty"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
			ArcPay        int `json:"arc_pay"`
			PayFreeWatch  int `json:"pay_free_watch"`
		} `json:"rights"`
		Owner struct {
			Mid  int64  `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid      int `json:"aid"`
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Favorite int `json:"favorite"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Dislike  int `json:"dislike"`
			Vt       int `json:"vt"`
			Vv       int `json:"vv"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		ShortLinkV2 string `json:"short_link_v2"`
		FirstFrame  string `json:"first_frame"`
		PubLocation string `json:"pub_location,omitempty"`
		Bvid        string `json:"bvid"`
		Score       int    `json:"score"`
		Others      []struct {
			Aid       int    `json:"aid"`
			Videos    int    `json:"videos"`
			Tid       int    `json:"tid"`
			Tname     string `json:"tname"`
			Copyright int    `json:"copyright"`
			Pic       string `json:"pic"`
			Title     string `json:"title"`
			Pubdate   int    `json:"pubdate"`
			Ctime     int    `json:"ctime"`
			Desc      string `json:"desc"`
			State     int    `json:"state"`
			Attribute int    `json:"attribute"`
			Duration  int    `json:"duration"`
			MissionId int    `json:"mission_id,omitempty"`
			Rights    struct {
				Bp            int `json:"bp"`
				Elec          int `json:"elec"`
				Download      int `json:"download"`
				Movie         int `json:"movie"`
				Pay           int `json:"pay"`
				Hd5           int `json:"hd5"`
				NoReprint     int `json:"no_reprint"`
				Autoplay      int `json:"autoplay"`
				UgcPay        int `json:"ugc_pay"`
				IsCooperation int `json:"is_cooperation"`
				UgcPayPreview int `json:"ugc_pay_preview"`
				NoBackground  int `json:"no_background"`
				ArcPay        int `json:"arc_pay"`
				PayFreeWatch  int `json:"pay_free_watch"`
			} `json:"rights"`
			Owner struct {
				Mid  int64  `json:"mid"`
				Name string `json:"name"`
				Face string `json:"face"`
			} `json:"owner"`
			Stat struct {
				Aid      int `json:"aid"`
				View     int `json:"view"`
				Danmaku  int `json:"danmaku"`
				Reply    int `json:"reply"`
				Favorite int `json:"favorite"`
				Coin     int `json:"coin"`
				Share    int `json:"share"`
				NowRank  int `json:"now_rank"`
				HisRank  int `json:"his_rank"`
				Like     int `json:"like"`
				Dislike  int `json:"dislike"`
				Vt       int `json:"vt"`
				Vv       int `json:"vv"`
			} `json:"stat"`
			Dynamic   string `json:"dynamic"`
			Cid       int    `json:"cid"`
			Dimension struct {
				Width  int `json:"width"`
				Height int `json:"height"`
				Rotate int `json:"rotate"`
			} `json:"dimension"`
			AttributeV2 int    `json:"attribute_v2"`
			ShortLinkV2 string `json:"short_link_v2"`
			FirstFrame  string `json:"first_frame"`
			PubLocation string `json:"pub_location,omitempty"`
			Bvid        string `json:"bvid"`
			Score       int    `json:"score"`
			EnableVt    int    `json:"enable_vt"`
			SeasonId    int    `json:"season_id,omitempty"`
			UpFromV2    int    `json:"up_from_v2,omitempty"`
		} `json:"others,omitempty"`
		EnableVt int `json:"enable_vt"`
		SeasonId int `json:"season_id,omitempty"`
		UpFromV2 int `json:"up_from_v2,omitempty"`
	} `json:"list"`
}

type VideoRankingDataVO struct {
	Aid      int    `json:"aid"`
	Bvid     string `json:"bvid"`
	Cid      int    `json:"cid"`
	Pic      string `json:"pic"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Duration string `json:"duration"` //稿件总时长(所有分P)	单位为秒
	View     string `json:"view"`     //播放数
	Reply    string `json:"reply"`    //评论数
	Like     string `json:"like"`     //点赞数
}

func (d *VideoRankingData) ToVO(picPrefix string) []*VideoRankingDataVO {
	var vo []*VideoRankingDataVO
	for _, v := range d.List {

		u, err := url.Parse(v.Pic)
		if err != nil {
			continue
		}

		vo = append(vo, &VideoRankingDataVO{
			Aid:      v.Aid,
			Bvid:     v.Bvid,
			Cid:      v.Cid,
			Pic:      picPrefix + filepath.Base(u.Path),
			Title:    v.Title,
			Desc:     v.Desc,
			Duration: toHM(v.Duration),
			View:     toHanz(v.Stat.View),
			Reply:    toHanz(v.Stat.Reply),
			Like:     toHanz(v.Stat.Like),
		})
	}
	return vo
}

func toHM(num int) string {
	t := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
	d := t.Add(time.Duration(num) * time.Second)
	str := d.Format("15:04:05")
	if str[:2] == "00" {
		return str[3:]
	}
	return str
}

func toHanz(num int) string {
	if num < 1000 { //999
		return cast.ToString(num)
	} else if num < 10000 { //9k
		return fmt.Sprintf("%.1fk", float64(num)/1000)
	} else if num < 100000000 { //999w
		return fmt.Sprintf("%.1fw", float64(num)/10000)
	} else {
		return fmt.Sprintf("%.1fe", float64(num)/100000000)
	}
}
