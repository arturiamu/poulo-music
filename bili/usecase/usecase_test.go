package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	ua "github.com/wux1an/fake-useragent"
	"poulo-music/config"
	"poulo-music/models"
	"testing"
)

var mockDir = &config.Dir{
	BaseDir: ".",
}

func TestGetVideoViewByBVID(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	data, err := uc.GetVideoViewByBvid(context.Background(), "BV1H3411C7ZJ")
	//data, err := uc.GetVideoViewByBvid("BV1xxxxxxxxx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", data)
}

func TestGetVideoViewDetailByBVID(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	data, err := uc.GetVideoViewByBvid(context.Background(), "BV1H3411C7ZJ")
	//data, err := uc.VideoViewDetailByBVID("BV1xxxxxxxxx")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}

func TestVideoPlayerWbiPlayUrl(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	data, err := uc.GetVideoPlayerWbiPlayUrl(context.Background(), "BV1H3411C7ZJ", 426064004)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", data)
}

func TestGetSearchAllResult(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	data, err := uc.GetSearchAllResult(context.Background(), "周杰伦")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}

func TestUseCase_GetSearchTypeResult(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	data, err := uc.GetSearchTypeResult(context.Background(),
		&models.SearchTypeParam{
			SearchType: "video",
			Keyword:    "周杰伦",
			Page:       1,
		}, false, "./cache")
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := json.Marshal(data)

	fmt.Println(string(bytes))

	fmt.Println(data.ToVO(""))
}

func TestUseCase_GetVideoRanking(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	_, err := uc.GetVideoRanking(context.Background(), true, "./cache")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Test_UA(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(ua.Random())
	}
}

func TestUseCase_GetVideoMp4ByUri(t *testing.T) {
	uc := NewUseCase(new(logrus.Logger), mockDir)
	err := uc.GetVideoMp4ByUri(context.Background(), "BV1H3411C7ZJ", "./BV1H3411C7ZJ.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestUseCase_GetVideoMp4ByUri2(t *testing.T) {
	var uri = "https://cn-lnsy-cu-01-04.bilivideo.com/upgcxcode/84/72/1451667284/1451667284-1-16.mp4?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1709289966&gen=playurlv2&os=bcache&oi=1875637252&trid=0000e3b0193a4920495eae7ddd306b386c61u&mid=0&platform=pc&upsig=7bd3d94c38e17dd2ec5ae99a343a3378&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform&cdnid=3283&bvc=vod&nettype=0&orderid=0,3&buvid=&build=0&f=u_0_0&agrr=1&bw=38032&logo=80000000"
	uc := NewUseCase(new(logrus.Logger), mockDir)
	err := uc.GetVideoMp4ByUri(context.Background(), uri, "./BV1H3411C7ZJ_1.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
}
