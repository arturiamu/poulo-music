package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/url"
	"os"
	"path/filepath"
	"poulo-music/bili"
	"poulo-music/config"
	"poulo-music/models"
	"strings"
	"sync"
)

var _ bili.UseCase = (*UseCase)(nil)

const (
	Sep = string(os.PathSeparator)
)

type UseCase struct {
	log *logrus.Logger
	dir *config.Dir
}

func NewUseCase(log *logrus.Logger, dir *config.Dir) *UseCase {
	return &UseCase{log: log, dir: dir}
}

func (uc *UseCase) getErr(code int) error {
	switch code {
	case 0:
		return nil
	case -400:
		return bili.ErrRequestIsIncorrect
	case -403:
		return bili.ErrInsufficientPermissions
	case -404:
		return bili.ErrNoVideo
	case -412:
		return bili.ErrSearchBlocked
	case 62002:
		return bili.ErrManuscriptIsNotVisible
	case 62004:
		return bili.ErrManuscriptIsUnderReview
	default:
		return bili.ErrUnknownError
	}
}

func (uc *UseCase) GetVideoViewByBvid(ctx context.Context, bvid string) (*models.VideoViewData, error) {
	return uc.getVideoViewByBvidOrAid(ctx, bvid, 0)
}

func (uc *UseCase) GetVideoViewByAid(ctx context.Context, aid int) (*models.VideoViewData, error) {
	return uc.getVideoViewByBvidOrAid(ctx, "", aid)
}

// getVideoViewByBvidOrAid
// 获取视频详细信息(web端)
// https://api.bilibili.com/x/web-interface/view
// 请求方式：GET
// 认证方式：Cookie(SESSDATA)
// 限制游客访问的视频需要登录
// url参数：
// 参数名	类型	内容	必要性	备注
// aid	num	稿件avid	必要(可选)	avid与bvid任选一个
// bvid	str	稿件bvid	必要(可选)	avid与bvid任选一个
// json回复：见 models
func (uc *UseCase) getVideoViewByBvidOrAid(ctx context.Context, bvid string, aid int) (*models.VideoViewData, error) {
	if bvid == "" && aid == 0 {
		return nil, bili.ErrInvalidParameter
	}

	var url = "https://api.bilibili.com/x/web-interface/view"

	bytes, err := NewHttp().SetUrl(url).AddParams(map[string]string{"bvid": bvid, "aid": fmt.Sprintf("%d", aid)}).Do()
	if err != nil {
		return nil, err
	}

	var videoView models.VideoViewResp

	err = json.Unmarshal(bytes, &videoView)
	if err != nil {
		return nil, err
	}

	if err := uc.getErr(videoView.Code); err != nil {
		return nil, err
	}

	return &videoView.Data, nil
}

func (uc *UseCase) GetVideoViewDetailByBvid(ctx context.Context, bvid string) (*models.VideoViewDetailData, error) {
	return uc.getVideoViewDetailByBvidOrAid(ctx, bvid, 0)
}

func (uc *UseCase) GetVideoViewDetailByAid(ctx context.Context, aid int) (*models.VideoViewDetailData, error) {
	return uc.getVideoViewDetailByBvidOrAid(ctx, "", aid)
}

// getVideoViewDetailByBvidOrAid
// 获取视频超详细信息(web端)
// https://api.bilibili.com/x/web-interface/view/detail
// https://api.bilibili.com/x/web-interface/wbi/view/detail
// 请求方式：GET
// 认证方式：Cookie(SESSDATA)
// 鉴权方式：Wbi 签名
// 限制游客访问的视频需要登录
// url参数：
// 参数名	类型	内容	必要性	备注
// aid	num	稿件avid	必要(可选)	avid与bvid任选一个
// bvid	str	稿件bvid	必要(可选)	avid与bvid任选一个
// json回复：
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
func (uc *UseCase) getVideoViewDetailByBvidOrAid(ctx context.Context, bvid string, aid int) (*models.VideoViewDetailData, error) {
	if bvid == "" && aid == 0 {
		return nil, bili.ErrInvalidParameter
	}

	var url = "https://api.bilibili.com/x/web-interface/view/detail"

	bytes, err := NewHttp().SetUrl(url).AddParams(map[string]string{"bvid": bvid, "aid": fmt.Sprintf("%d", aid)}).Do()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	var videoViewDetail models.VideoViewDetailResp

	err = json.Unmarshal(bytes, &videoViewDetail)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if err := uc.getErr(videoViewDetail.Code); err != nil {
		uc.log.Error(err)
		return nil, err
	}

	return &videoViewDetail.Data, nil
}

// GetVideoMp4ByUri
// uri 可以是 bvid 或者 url
// savePath 绝对路径,完整文件名
func (uc *UseCase) GetVideoMp4ByUri(ctx context.Context, uri string, filename string) (err error) {
	if strings.HasPrefix(uri, "http://") || strings.HasPrefix(uri, "https://") {
		//return uc.GetVideoMp4ByUrl(ctx, uri, savePath)
		return uc.SaveFile(uri, filename)
	} else {
		return uc.GetVideoMp4ByBvid(ctx, uri, filename)
	}
}

// GetVideoMp4ByBvid
// savePath 绝对路径,完整文件名
func (uc *UseCase) GetVideoMp4ByBvid(ctx context.Context, bvid string, filename string) error {
	videoViewDetail, err := uc.GetVideoViewDetailByBvid(ctx, bvid)
	if err != nil {
		uc.log.Error(err)
		return err
	}

	videoPlayerWbiPlayUrl, err := uc.GetVideoPlayerWbiPlayUrl(ctx, bvid, videoViewDetail.View.Cid)
	if err != nil {
		uc.log.Error(err)
		return err
	}

	if len(videoPlayerWbiPlayUrl.Durl) == 0 {
		return bili.ErrNoMp4Video
	}
	return uc.SaveFile(videoPlayerWbiPlayUrl.Durl[0].Url, filename)
}

func (uc *UseCase) GetVideoPlayerWbiPlayUrl(ctx context.Context, bvid string, cid int) (*models.VideoPlayerWbiPlayUrlData, error) {

	if bvid == "" || cid == 0 {
		return nil, bili.ErrInvalidParameter
	}

	var url = "https://api.bilibili.com/x/player/playurl"

	bytes, err := NewHttp().SetUrl(url).AddParams(map[string]string{"bvid": bvid, "cid": fmt.Sprintf("%d", cid)}).Do()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	var videoPlayer models.VideoPlayerWbiPlayUrlResp

	err = json.Unmarshal(bytes, &videoPlayer)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if err := uc.getErr(videoPlayer.Code); err != nil {
		uc.log.Error(err)
		return nil, err
	}

	return &videoPlayer.Data, nil
}

func (uc *UseCase) getTempCookie() (string, error) {
	var url = "https://www.bilibili.com/"
	response, err := NewHttp().SetUrl(url).DoDetail()
	if err != nil {
		uc.log.Error(err)
		return "", err
	}

	var cookieStr string
	for _, cookie := range response.Cookies() {
		c := fmt.Sprintf("%s=%s; ", cookie.Name, cookie.Value)
		cookieStr += c
	}

	return cookieStr, nil
}

// GetSearchAllResult
// 综合搜索（web端）
// https://api.bilibili.com/x/web-interface/wbi/search/all/v2
// https://api.bilibili.com/x/web-interface/search/all/v2 （旧链接 x）
// 方式：GET
// 认证方式：Cookie（SESSDATA）
// 鉴权方式：Wbi 签名
// 返回和关键字相关的20条信息
// 综合搜索为默认搜索方式，主要用于优先搜索用户、影视、番剧、游戏、话题等，并加载第一页的20项相关视频，还用于展示各个类型的结果数目，便于进一步分类搜索
func (uc *UseCase) GetSearchAllResult(ctx context.Context, keyword string) (*models.SearchAllData, error) {
	var url = "https://api.bilibili.com/x/web-interface/wbi/search/all/v2"
	cookie, err := uc.getTempCookie()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	bytes, err := NewHttp().SetUrl(url).AddParams(map[string]string{"keyword": keyword}).SetHeader("cookie", cookie).Do()

	var searchAll models.SearchAllResp
	err = json.Unmarshal(bytes, &searchAll)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if err := uc.getErr(searchAll.Code); err != nil {
		uc.log.Error(err)
		return nil, err
	}

	return &searchAll.Data, nil
}

// GetSearchTypeResult
// 分类搜索（web端）
// https://api.bilibili.com/x/web-interface/wbi/search/type
// https://api.bilibili.com/x/web-interface/search/type （旧链接 x）
// 请求方式：GET
// 认证方式：Cookie（SESSDATA）
// 鉴权方式：Wbi 签名
// 根据关键词进行搜索，返回结果每页20项
func (uc *UseCase) GetSearchTypeResult(ctx context.Context, param *models.SearchTypeParam, saveFile bool, saveDir string) (*models.SearchTypeData, error) {
	var searchUrl = "https://api.bilibili.com/x/web-interface/wbi/search/type"
	cookie, err := uc.getTempCookie()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	bytes, err := NewHttp().SetUrl(searchUrl).SetHeader("cookie", cookie).AddParams(param.ToHttpParamsMap()).Do()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}
	var searchType models.SearchTypeResp
	err = json.Unmarshal(bytes, &searchType)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if err := uc.getErr(searchType.Code); err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if saveFile {
		uc.saveSearchTypeFile(&searchType.Data, saveDir)
	}

	return &searchType.Data, nil
}

func (uc *UseCase) saveSearchTypeFile(searchTypeData *models.SearchTypeData, saveDir string) {
	var wg sync.WaitGroup
	wg.Add(len(searchTypeData.Result))
	for _, s := range searchTypeData.Result {
		u, err := url.Parse(s.Pic)
		if err != nil {
			uc.log.Error(err)
			continue
		}
		go func(url, filename string) {
			defer wg.Done()
			if !strings.HasPrefix(url, "http") {
				url = "http:" + url
			}
			uc.SaveFile(url, filename)
		}(s.Pic, saveDir+Sep+filepath.Base(u.Path))
	}
	wg.Wait()
}

// GetVideoRanking
// 获取分区视频排行榜列表 (音乐类型)
// https://api.bilibili.com/x/web-interface/ranking/v2
// 请求方式：GET
// 获取稿件内容质量，近期的数据前100个稿件，动态更新。
// url参数：
// 参数名	类型	内容	必要性	备注
// rid	num	目标分区tid	非必要	可参考视频分区一览，只支持主分区
// type	str	未知	非必要	默认为：all，且为目前唯一已知值。怀疑为稿件类型，但没有找到其他值佐证。
func (uc *UseCase) GetVideoRanking(ctx context.Context, saveFile bool, saveDir string) (*models.VideoRankingData, error) {
	var rankingUrl = "https://api.bilibili.com/x/web-interface/ranking/v2"

	bytes, err := NewHttp().SetUrl(rankingUrl).AddParam("rid", "3").Do()
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	var rankingResp models.VideoRankingResp
	err = json.Unmarshal(bytes, &rankingResp)
	if err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if err := uc.getErr(rankingResp.Code); err != nil {
		uc.log.Error(err)
		return nil, err
	}

	if saveFile {
		uc.saveRankingFile(&rankingResp.Data, saveDir)
	}

	return &rankingResp.Data, nil
}

// ~/Library/Caches/BiliMusicify/cache/pic.jpg
func (uc *UseCase) saveRankingFile(videoRankingData *models.VideoRankingData, saveDir string) {
	var wg sync.WaitGroup
	wg.Add(len(videoRankingData.List))
	for _, s := range videoRankingData.List {
		u, err := url.Parse(s.Pic)
		if err != nil {
			uc.log.Error(err)
			continue
		}
		go func(url, filename string) {
			defer wg.Done()
			if !strings.HasPrefix(url, "http") {
				url = "http:" + url
			}
			uc.SaveFile(url, filename)
		}(s.Pic, saveDir+Sep+filepath.Base(u.Path))
	}
	wg.Wait()
}

func (uc *UseCase) SaveFile(url, filename string) error {
	bytes, err := NewHttp().SetUrl(url).Do()
	if err != nil {
		uc.log.Error(err)
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		uc.log.Error(err)
		return err
	}

	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		uc.log.Error(err)
		return err
	}

	return nil
}
