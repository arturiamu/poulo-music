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

type Server struct {
	platformMap map[models.Platform]platform.Platform
	codec       codec.UseCase
	ctx         context.Context
	cfg         *config.Config
}

func NewServer(cfg *config.Config, log *logrus.Logger) *Server {
	_, err := initDB(cfg.Dir)
	if err != nil {
		panic(err)
	}

	biliClient := bili.NewBili(log, &cfg.Dir)
	qqClient := qq.NewQQ(log)
	miguClient := migu.NewMigu(log)

	return &Server{
		cfg: cfg,
		platformMap: map[models.Platform]platform.Platform{
			models.PlatformBili: biliClient,
			models.PlatformQQ:   qqClient,
			models.PlatformMigu: miguClient,
		},
	}
}

func initDB(dir config.Dir) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s%s%s", dir.BaseDir, string(filepath.Separator), "data.db")), &gorm.Config{})
}

func (s *Server) SetContext(ctx context.Context) {
	s.ctx = ctx
}

func (s *Server) RunServer(ctx context.Context) {
	s.ctx = ctx

	fs := http.FileServer(http.Dir(s.cfg.Dir.BaseDir + string(filepath.Separator) + "cache"))

	http.Handle("/", middleware.CORS(fs))

	http.ListenAndServe(":51730", nil)
}

func (s *Server) AggregateSearch(platform models.Platform, param models.GetSearchParam) (data []models.GetSearchResp, err error) {
	return s.platformMap[platform].GetSearch(s.ctx, param)
}

func (s *Server) AggregateHotContent(platform models.Platform, param models.GetHotContentParam) (data []models.GetHotContentResp, err error) {
	return s.platformMap[platform].GetHotContent(s.ctx, param)
}

func (s *Server) AggregateMusic(platform models.Platform, param models.GetMusicParam) (data models.Music, err error) {
	return s.platformMap[platform].GetMusic(s.ctx, param)
}
