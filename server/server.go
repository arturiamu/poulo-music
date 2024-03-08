package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"poulo-music/codec"
	"poulo-music/config"
	"poulo-music/middleware"
	"poulo-music/models"
	"poulo-music/platform"
	"poulo-music/platform/bili"
	"poulo-music/platform/migu"
	"poulo-music/platform/qq"
)

const (
	Sep = string(filepath.Separator)
)

type Server struct {
	platformMap map[models.Platform]platform.Platform
	codec       codec.UseCase
	ctx         context.Context
	cfg         *config.Config
}

func NewServer(cfg *config.Config, log *logrus.Logger) (*Server, error) {
	_, err := initDB(cfg.BaseDir + Sep + "data")
	if err != nil {
		return nil, err
	}

	var cacheDir = cfg.BaseDir + Sep + "cache"
	biliClient := bili.NewBili(log, cacheDir)
	qqClient := qq.NewQQ(log, cacheDir)
	miguClient := migu.NewMigu(log, cacheDir)

	var svr = &Server{
		cfg: cfg,
		platformMap: map[models.Platform]platform.Platform{
			models.PlatformBili: biliClient,
			models.PlatformQQ:   qqClient,
			models.PlatformMigu: miguClient,
		},
	}

	return svr, nil
}

func initDB(dataPath string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s%s%s", dataPath, string(filepath.Separator), "data.db")), &gorm.Config{})
}

func (s *Server) RunServer(ctx context.Context) {
	s.ctx = ctx

	fs := http.FileServer(http.Dir(s.cfg.BaseDir + Sep + "cache"))

	http.Handle("/", middleware.CORS(fs))

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.FileServerPort), nil)
	if err != nil {
		panic(err)
	}
}

func (s *Server) AggregateSearch(platform models.Platform, param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	if platform == models.PlatformAll {
		return s.aggregateSearchAll(platform, param)
	}
	return s.platformMap[platform].GetSearch(s.ctx, param)
}

func (s *Server) aggregateSearchAll(platform models.Platform, param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	return
}

func (s *Server) AggregateHotContent(platform models.Platform, param models.GetHotContentParam) (data []models.GetHotContentResp, err error) {
	return s.platformMap[platform].GetHotContent(s.ctx, param)
}

func (s *Server) AggregateMusic(platform models.Platform, param models.GetMusicParam) (data models.Music, err error) {
	return s.platformMap[platform].GetMusic(s.ctx, param)
}
