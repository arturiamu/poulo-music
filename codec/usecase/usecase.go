package usecase

import (
	"fmt"
	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"poulo-music/codec"
)

var _ codec.UseCase = (*UseCase)(nil)

type UseCase struct {
	log *logrus.Logger
}

func NewUseCase(log *logrus.Logger) *UseCase {
	return &UseCase{log: log}
}

func (u UseCase) GetAudioFromUrl(url, output string) error {
	err := ffmpeg.Input(url, ffmpeg.KwArgs{"user_agent": fmt.Sprintf("\"%s\"", browser.Random()), "headers": "referer: https://www.bilibili.com"}).
		Output(output, ffmpeg.KwArgs{"vn": "", "acodec": "libmp3lame", "qscale:a": 0}).
		Run()
	if err != nil {
		u.log.Error(err)
		return err
	}
	return err
}

func (u UseCase) GetAudioFromFile(input, output string) error {
	// ffmpeg -i video.mp4 -vn -acodec libmp3lame -qscale:a 2 output.mp3
	err := ffmpeg.Input(input).Output(output, ffmpeg.KwArgs{
		"vn": "", "acodec": "libmp3lame", "qscale:a": 0,
	}).Run()
	if err != nil {
		u.log.Error(err)
		return err
	}
	return err
}
