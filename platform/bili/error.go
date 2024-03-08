package bili

import "errors"

var (
	ErrInvalidParameter        = errors.New("invalid parameter")
	ErrRequestIsIncorrect      = errors.New("request is incorrect")       //-400：请求错误
	ErrInsufficientPermissions = errors.New("insufficient permissions")   //-403：权限不足
	ErrNoVideo                 = errors.New("no video")                   //-404：无视频
	ErrManuscriptIsNotVisible  = errors.New("manuscript not visible")     //62002：稿件不可见
	ErrManuscriptIsUnderReview = errors.New("manuscript is under review") //62004：稿件审核中
	ErrSearchBlocked           = errors.New("search blocked")             //-412：搜索被拦截

	ErrNoMp4Video   = errors.New("no mp4 video")  //无mp4视频地址
	ErrUnknownError = errors.New("unknown error") //未知错误
)
