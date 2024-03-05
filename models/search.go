package models

import (
	"github.com/spf13/cast"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	Sep string = string(filepath.Separator)
)

type SearchAllResp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Ttl     int           `json:"ttl"`
	Data    SearchAllData `json:"data"`
}

type SearchAllData struct {
	Seid           string `json:"seid"`
	Page           int    `json:"page"`
	Pagesize       int    `json:"pagesize"`
	NumResults     int    `json:"numResults"`
	NumPages       int    `json:"numPages"`
	SuggestKeyword string `json:"suggest_keyword"`
	RqtType        string `json:"rqt_type"`
	CostTime       struct {
		Total               string `json:"total"`
		FetchLexicon        string `json:"fetch_lexicon"`
		ParamsCheck         string `json:"params_check"`
		IsRiskQuery         string `json:"is_risk_query"`
		IllegalHandler      string `json:"illegal_handler"`
		MainHandler         string `json:"main_handler"`
		GetUpuserLiveStatus string `json:"get upuser live status"`
		MysqlRequest        string `json:"mysql_request"`
		AsRequestFormat     string `json:"as_request_format"`
		AsRequest           string `json:"as_request"`
		DeserializeResponse string `json:"deserialize_response"`
		AsResponseFormat    string `json:"as_response_format"`
	} `json:"cost_time"`
	ExpList struct {
		Field1 bool `json:"150000"`
	} `json:"exp_list"`
	EggHit   int `json:"egg_hit"`
	Pageinfo struct {
		Video struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"video"`
		Bangumi struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"bangumi"`
		Special struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"special"`
		Topic struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"topic"`
		Upuser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"upuser"`
		Tv struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"tv"`
		Movie struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"movie"`
		MediaBangumi struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"media_bangumi"`
		MediaFt struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"media_ft"`
		RelatedSearch struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"related_search"`
		User struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"user"`
		Activity struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"activity"`
		OperationCard struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"operation_card"`
		Pgc struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"pgc"`
		Live struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live"`
		LiveAll struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_all"`
		LiveMaster struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_master"`
		Article struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"article"`
		Intervene struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"intervene"`
		LiveRoom struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_room"`
		BiliUser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"bili_user"`
		LiveUser struct {
			Total      int `json:"total"`
			NumResults int `json:"numResults"`
			Pages      int `json:"pages"`
		} `json:"live_user"`
	} `json:"pageinfo"`
	TopTlist struct {
		Video         int `json:"video"`
		Bangumi       int `json:"bangumi"`
		Special       int `json:"special"`
		Topic         int `json:"topic"`
		Upuser        int `json:"upuser"`
		Tv            int `json:"tv"`
		Movie         int `json:"movie"`
		Card          int `json:"card"`
		MediaBangumi  int `json:"media_bangumi"`
		MediaFt       int `json:"media_ft"`
		Pgc           int `json:"pgc"`
		Live          int `json:"live"`
		User          int `json:"user"`
		Activity      int `json:"activity"`
		OperationCard int `json:"operation_card"`
		LiveMaster    int `json:"live_master"`
		Article       int `json:"article"`
		Intervene     int `json:"intervene"`
		LiveRoom      int `json:"live_room"`
		BiliUser      int `json:"bili_user"`
		LiveUser      int `json:"live_user"`
	} `json:"top_tlist"`
	ShowColumn       int      `json:"show_column"`
	ShowModuleList   []string `json:"show_module_list"`
	AppDisplayOption struct {
		IsSearchPageGrayed int `json:"is_search_page_grayed"`
	} `json:"app_display_option"`
	InBlackKey int `json:"in_black_key"`
	InWhiteKey int `json:"in_white_key"`
	Result     []struct {
		ResultType string `json:"result_type"`
		Data       []struct {
			Type        string `json:"type"`
			Mid         int64  `json:"mid"`
			Uname       string `json:"uname,omitempty"`
			Usign       string `json:"usign,omitempty"`
			Fans        int    `json:"fans,omitempty"`
			Videos      int    `json:"videos,omitempty"`
			Upic        string `json:"upic"`
			FaceNft     int    `json:"face_nft,omitempty"`
			FaceNftType int    `json:"face_nft_type,omitempty"`
			VerifyInfo  string `json:"verify_info,omitempty"`
			Level       int    `json:"level,omitempty"`
			Gender      int    `json:"gender,omitempty"`
			IsUpuser    int    `json:"is_upuser,omitempty"`
			IsLive      int    `json:"is_live,omitempty"`
			RoomId      int    `json:"room_id,omitempty"`
			Res         []struct {
				Aid           int    `json:"aid"`
				Bvid          string `json:"bvid"`
				Title         string `json:"title"`
				Pubdate       int    `json:"pubdate"`
				Arcurl        string `json:"arcurl"`
				Pic           string `json:"pic"`
				Play          string `json:"play"`
				Dm            int    `json:"dm"`
				Coin          int    `json:"coin"`
				Fav           int    `json:"fav"`
				Desc          string `json:"desc"`
				Duration      string `json:"duration"`
				IsPay         int    `json:"is_pay"`
				IsUnionVideo  int    `json:"is_union_video"`
				IsChargeVideo int    `json:"is_charge_video"`
				Vt            int    `json:"vt"`
				EnableVt      int    `json:"enable_vt"`
				VtDisplay     string `json:"vt_display"`
			} `json:"res,omitempty"`
			OfficialVerify struct {
				Type int    `json:"type"`
				Desc string `json:"desc"`
			} `json:"official_verify,omitempty"`
			HitColumns     []string `json:"hit_columns"`
			IsSeniorMember int      `json:"is_senior_member,omitempty"`
			Expand         struct {
				IsPowerUp    bool        `json:"is_power_up"`
				SystemNotice interface{} `json:"system_notice"`
			} `json:"expand,omitempty"`
			Id               int           `json:"id,omitempty"`
			Author           string        `json:"author,omitempty"`
			Typeid           string        `json:"typeid,omitempty"`
			Typename         string        `json:"typename,omitempty"`
			Arcurl           string        `json:"arcurl,omitempty"`
			Aid              int           `json:"aid,omitempty"`
			Bvid             string        `json:"bvid,omitempty"`
			Title            string        `json:"title,omitempty"`
			Description      string        `json:"description,omitempty"`
			Arcrank          string        `json:"arcrank,omitempty"`
			Pic              string        `json:"pic,omitempty"`
			Play             int           `json:"play,omitempty"`
			VideoReview      int           `json:"video_review,omitempty"`
			Favorites        int           `json:"favorites,omitempty"`
			Tag              string        `json:"tag,omitempty"`
			Review           int           `json:"review,omitempty"`
			Pubdate          int           `json:"pubdate,omitempty"`
			Senddate         int           `json:"senddate,omitempty"`
			Duration         string        `json:"duration,omitempty"`
			Badgepay         bool          `json:"badgepay,omitempty"`
			ViewType         string        `json:"view_type,omitempty"`
			IsPay            int           `json:"is_pay,omitempty"`
			IsUnionVideo     int           `json:"is_union_video,omitempty"`
			RecTags          interface{}   `json:"rec_tags"`
			NewRecTags       []interface{} `json:"new_rec_tags,omitempty"`
			RankScore        int           `json:"rank_score,omitempty"`
			Like             int           `json:"like,omitempty"`
			Corner           string        `json:"corner,omitempty"`
			Cover            string        `json:"cover,omitempty"`
			Desc             string        `json:"desc,omitempty"`
			Url              string        `json:"url,omitempty"`
			RecReason        string        `json:"rec_reason,omitempty"`
			Danmaku          int           `json:"danmaku,omitempty"`
			BizData          interface{}   `json:"biz_data"`
			IsChargeVideo    int           `json:"is_charge_video,omitempty"`
			Vt               int           `json:"vt,omitempty"`
			EnableVt         int           `json:"enable_vt,omitempty"`
			VtDisplay        string        `json:"vt_display,omitempty"`
			Subtitle         string        `json:"subtitle,omitempty"`
			EpisodeCountText string        `json:"episode_count_text,omitempty"`
			ReleaseStatus    int           `json:"release_status,omitempty"`
			IsIntervene      int           `json:"is_intervene,omitempty"`
		} `json:"data"`
	} `json:"result"`
	IsSearchPageGrayed int `json:"is_search_page_grayed"`
}

type SearchTypeParam struct {
	//视频：video
	//番剧：media_bangumi
	//影视：media_ft
	//直播间及主播：live
	//直播间：live_room
	//主播：live_user
	//专栏：article
	//话题：topic
	//用户：bili_user
	//相簿：photo
	SearchType string `json:"search_type"` //search_type	string // str	搜索目标类型	必要

	Keyword string `json:"keyword"` //keyword	str	需要搜索的关键词	必要

	//搜索类型为视频、专栏及相簿时：
	//默认为totalrank
	//综合排序：totalrank
	//最多点击：click
	//最新发布：pubdate
	//最多弹幕：dm
	//最多收藏：stow
	//最多评论：scores
	//最多喜欢：attention（仅用于专栏）
	//----------------------------
	//搜索结果为直播间时：
	//默认为online
	//人气直播：online
	//最新开播：live_time
	//----------------------------
	//搜索结果为用户时：
	//默认为0
	//默认排序：0
	//粉丝数：fans
	//用户等级：level
	Order string `json:"order"` //order	str	结果排序方式	非必要

	//仅用于搜索用户
	//默认为0
	//由高到低：0
	//由低到高：1
	OrderSort int `json:"order_sort"` //order_sort	num	用户粉丝数及等级排序顺序	非必要

	//仅用于搜索用户
	//默认为0
	//全部用户：0
	//up主：1
	//普通用户：2
	//认证用户：3
	UserType int `json:"user_type"` //user_type	num	用户分类筛选	非必要

	//仅用于搜索视频
	//默认为0
	//全部时长：0
	//10分钟以下：1
	//10-30分钟：2
	//30-60分钟：3
	//60分钟以上：4
	Duration int `json:"duration"` //duration	num	视频时长筛选	非必要

	//仅用于搜索视频
	//默认为0
	//全部分区：0
	//筛选分区：目标分区tid
	Tids int `json:"tids"` //tids	num	视频分区筛选	非必要

	//搜索结果为专栏时：
	//默认为0
	//全部分区：0
	//动画：2
	//游戏：1
	//影视：28
	//生活：3
	//兴趣：29
	//轻小说：16
	//科技：17
	//--------
	//搜索结果为相簿时：
	//默认为0
	//全部分区：0
	//画友：1
	//摄影：2
	CategoryId int `json:"category_id"` //category_id	num	专栏及相簿分区筛选	非必要

	Page int `json:"page"` //page	num	页码	非必要	默认为1
}

func (p SearchTypeParam) ToHttpParamsMap() map[string]string {
	return map[string]string{
		"search_type": p.SearchType,
		"keyword":     p.Keyword,
		"order":       p.Order,
		"order_sort":  cast.ToString(p.OrderSort),
		"user_type":   cast.ToString(p.UserType),
		"duration":    cast.ToString(p.Duration),
		"tids":        cast.ToString(p.Tids),
		"category_id": cast.ToString(p.CategoryId),
		"page":        cast.ToString(p.Page),
	}
}

type SearchTypeResp struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Ttl     int            `json:"ttl"`
	Data    SearchTypeData `json:"data"`
}

type SearchTypeData struct {
	Seid           string `json:"seid"`
	Page           int    `json:"page"`
	Pagesize       int    `json:"pagesize"`
	NumResults     int    `json:"numResults"`
	NumPages       int    `json:"numPages"`
	SuggestKeyword string `json:"suggest_keyword"`
	RqtType        string `json:"rqt_type"`
	CostTime       struct {
		Total               string `json:"total"`
		FetchLexicon        string `json:"fetch_lexicon"`
		ParamsCheck         string `json:"params_check"`
		IsRiskQuery         string `json:"is_risk_query"`
		IllegalHandler      string `json:"illegal_handler"`
		MainHandler         string `json:"main_handler"`
		AsRequestFormat     string `json:"as_request_format"`
		AsRequest           string `json:"as_request"`
		DeserializeResponse string `json:"deserialize_response"`
		AsResponseFormat    string `json:"as_response_format"`
	} `json:"cost_time"`
	ExpList struct {
		Field1  bool `json:"5510"`
		Field2  bool `json:"6609"`
		Field3  bool `json:"7706"`
		Field4  bool `json:"9902"`
		Field5  bool `json:"9924"`
		Field6  bool `json:"9931"`
		Field7  bool `json:"9940"`
		Field8  bool `json:"9961"`
		Field9  bool `json:"100202"`
		Field10 bool `json:"100307"`
		Field11 bool `json:"100805"`
		Field12 bool `json:"101405"`
		Field13 bool `json:"101506"`
		Field14 bool `json:"101800"`
		Field15 bool `json:"101904"`
		Field16 bool `json:"102501"`
		Field17 bool `json:"102600"`
		Field18 bool `json:"102804"`
		Field19 bool `json:"102901"`
		Field20 bool `json:"103006"`
		Field21 bool `json:"103106"`
		Field22 bool `json:"103207"`
		Field23 bool `json:"103303"`
		Field24 bool `json:"103400"`
		Field25 bool `json:"103501"`
	} `json:"exp_list"`
	EggHit int `json:"egg_hit"`
	Result []struct {
		Type             string        `json:"type"`
		Id               int           `json:"id"`
		Author           string        `json:"author"`
		Mid              int           `json:"mid"`
		Typeid           string        `json:"typeid"`
		Typename         string        `json:"typename"`
		Arcurl           string        `json:"arcurl"`
		Aid              int           `json:"aid"`
		Bvid             string        `json:"bvid"`
		Title            string        `json:"title"`
		Description      string        `json:"description"`
		Arcrank          string        `json:"arcrank"`
		Pic              string        `json:"pic"`
		Play             int           `json:"play"`
		VideoReview      int           `json:"video_review"`
		Favorites        int           `json:"favorites"`
		Tag              string        `json:"tag"`
		Review           int           `json:"review"`
		Pubdate          int           `json:"pubdate"`
		Senddate         int           `json:"senddate"`
		Duration         string        `json:"duration"`
		Badgepay         bool          `json:"badgepay"`
		HitColumns       []string      `json:"hit_columns"`
		ViewType         string        `json:"view_type"`
		IsPay            int           `json:"is_pay"`
		IsUnionVideo     int           `json:"is_union_video"`
		RecTags          interface{}   `json:"rec_tags"`
		NewRecTags       []interface{} `json:"new_rec_tags"`
		RankScore        int           `json:"rank_score"`
		Like             int           `json:"like"`
		Upic             string        `json:"upic"`
		Corner           string        `json:"corner"`
		Cover            string        `json:"cover"`
		Desc             string        `json:"desc"`
		Url              string        `json:"url"`
		RecReason        string        `json:"rec_reason"`
		Danmaku          int           `json:"danmaku"`
		BizData          interface{}   `json:"biz_data"`
		IsChargeVideo    int           `json:"is_charge_video"`
		Vt               int           `json:"vt"`
		EnableVt         int           `json:"enable_vt"`
		VtDisplay        string        `json:"vt_display"`
		Subtitle         string        `json:"subtitle"`
		EpisodeCountText string        `json:"episode_count_text"`
		ReleaseStatus    int           `json:"release_status"`
		IsIntervene      int           `json:"is_intervene"`
	} `json:"result"`
	ShowColumn int `json:"show_column"`
	InBlackKey int `json:"in_black_key"`
	InWhiteKey int `json:"in_white_key"`
}

func (s *SearchTypeData) ToVO(picDir string) []*SearchTypeDataVO {
	var result []*SearchTypeDataVO

	re := regexp.MustCompile(`<em class="[a-zA-Z]+">(.+?)</em>`)

	for _, item := range s.Result {
		u, err := url.Parse(item.Pic)
		if err != nil {
			continue
		}

		kwd := ""
		match := re.FindStringSubmatch(item.Title)
		if len(match) > 1 {
			kwd = match[1]
		}

		title := re.ReplaceAllString(item.Title, kwd)
		title = strings.ReplaceAll(title, "&amp;", "&")
		result = append(result, &SearchTypeDataVO{
			Aid:      item.Aid,
			Bvid:     item.Bvid,
			Cid:      item.Id,
			Pic:      picDir + Sep + filepath.Base(u.Path),
			Title:    title,
			Desc:     item.Desc,
			Duration: item.Duration,
			View:     toHanz(item.Play),
			Reply:    toHanz(item.Review),
			Like:     toHanz(item.Like),
		})
	}
	return result
}

type SearchTypeDataVO struct {
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
