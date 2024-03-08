package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"poulo-music/config"
	"poulo-music/logger"
	"poulo-music/models"
	"poulo-music/server"
)

// App struct
type App struct {
	svr *server.Server
	log *logrus.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	log, err := logger.Init(cfg.BaseDir)
	if err != nil {
		panic(err)
	}

	svr, err := server.NewServer(cfg, log)
	if err != nil {
		panic(err)
	}

	return &App{svr: svr, log: log}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.svr.RunServer(ctx)
}

func (a *App) GetSearch(platform models.Platform, keyword string) (vo models.VO) {
	a.log.Infof("GetSearch platform: %v, keyword: %v", platform, keyword)
	var param = models.GetSearchParam{Platform: platform, Keyword: keyword}
	data, err := a.svr.AggregateSearch(platform, param)
	if err != nil {
		return vo.Fail(err.Error())
	}
	return vo.Success(data)
}

func (a *App) GetHotContent(platform models.Platform) (vo models.VO) {
	a.log.Infof("GetHotContent platform: %v", platform)
	var param = models.GetHotContentParam{Platform: platform}
	data, err := a.svr.AggregateHotContent(platform, param)
	if err != nil {
		return vo.Fail(err.Error())
	}
	return vo.Success(data)
}

func (a *App) GetMusic(platform models.Platform, identifier string) (vo models.VO) {
	a.log.Infof("GetMusic platform: %v,identifier: %v", platform, identifier)
	var param = models.GetMusicParam{Platform: platform, Identifier: identifier}
	data, err := a.svr.AggregateMusic(platform, param)
	if err != nil {
		return vo.Fail(err.Error())
	}
	return vo.Success(data)
}
