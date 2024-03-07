package platform

import "poulo-music/models"

type Platform interface {
	GetSearch(param models.GetSearchParam) (data []models.GetSearchResp, err error)
	GetHotContent(param models.GetHotContentParam) (data []models.GetHotContentResp, err error)
	GetMusic(param models.GetMusicParam) (data models.Music, err error)
}
