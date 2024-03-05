package usecase

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestUseCase_GetAudioFromUrl(t *testing.T) {
	var url = "https://cn-hncs-cu-01-06.bilivideo.com/upgcxcode/04/40/426064004/426064004_da2-1-16.mp4?e=ig8euxZM2rNcNbRVhwdVhwdlhWdVhwdVhoNvNC8BqJIzNbfqXBvEqxTEto8BTrNvN0GvT90W5JZMkX_YN0MvXg8gNEV4NC8xNEV4N03eN0B5tZlqNxTEto8BTrNvNeZVuJ10Kj_g2UB02J0mN0B5tZlqNCNEto8BTrNvNC7MTX502C8f2jmMQJ6mqF2fka1mqx6gqj0eN0B599M=&uipk=5&nbs=1&deadline=1709200611&gen=playurlv2&os=bcache&oi=1875637252&trid=00008dac0c1eeb374fe098b3475af7b29b17u&mid=0&platform=pc&upsig=a8fd35e1e89e29528394935fc02dfc93&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,mid,platform&cdnid=3295&bvc=vod&nettype=0&orderid=0,3&buvid=&build=0&f=u_0_0&agrr=1&bw=50582&logo=80000000"
	var uc = NewUseCase(&logrus.Logger{})
	err := uc.GetAudioFromUrl(url, "BV1H3411C7ZJ.mp3")
	if err != nil {
		t.Error(err)
	}
}
