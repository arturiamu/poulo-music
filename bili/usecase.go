package bili

import (
	"context"
	"poulo-music/models"
)

type UseCase interface {
	GetVideoViewByBvid(ctx context.Context, bvid string) (*models.VideoViewData, error)
	GetVideoViewByAid(ctx context.Context, aid int) (*models.VideoViewData, error)

	GetVideoViewDetailByBvid(ctx context.Context, bvid string) (*models.VideoViewDetailData, error)
	GetVideoViewDetailByAid(ctx context.Context, aid int) (*models.VideoViewDetailData, error)

	GetVideoPlayerWbiPlayUrl(ctx context.Context, bvid string, cid int) (*models.VideoPlayerWbiPlayUrlData, error)

	GetVideoMp4ByUri(ctx context.Context, uri string, filename string) error

	GetSearchAllResult(ctx context.Context, keyword string) (*models.SearchAllData, error)
	GetSearchTypeResult(ctx context.Context, param *models.SearchTypeParam, saveFile bool, saveDir string) (*models.SearchTypeData, error)

	GetVideoRanking(ctx context.Context, saveFile bool, saveDir string) (*models.VideoRankingData, error)
}
