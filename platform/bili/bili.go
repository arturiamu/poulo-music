package bili

import (
	"github.com/sirupsen/logrus"
	"poulo-music/config"
	"poulo-music/models"
	"poulo-music/platform"
)

var _ platform.Platform = (*Bili)(nil)

type Bili struct {
	log *logrus.Logger
	dir *config.Dir
}

func NewBili(log *logrus.Logger, dir *config.Dir) *Bili {
	return &Bili{log: log, dir: dir}
}

func (b Bili) GetSearch(param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (b Bili) GetHotContent(param models.GetHotContentParam) (data []models.GetHotContentResp, err error) {
	//TODO implement me
	panic("implement me")
}

func (b Bili) GetMusic(param models.GetMusicParam) (data models.Music, err error) {
	//TODO implement me
	panic("implement me")
}
