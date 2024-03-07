package platform

import (
	"context"
	"poulo-music/models"
)

type Platform interface {
	GetSearch(ctx context.Context, param models.GetSearchParam) (data []models.GetSearchResp, err error)
	GetHotContent(ctx context.Context, param models.GetHotContentParam) (data []models.GetHotContentResp, err error)
	GetMusic(ctx context.Context, param models.GetMusicParam) (data models.Music, err error)
}
