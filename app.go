package main

import (
	"context"
	"poulo-music/config"
	"poulo-music/logger"
	"poulo-music/server"
)

// App struct
type App struct {
	svr *server.Server
}

// NewApp creates a new App application struct
func NewApp() *App {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	log, err := logger.Init(cfg.Dir.BaseDir)
	if err != nil {
		panic(err)
	}

	return &App{
		svr: server.NewServer(cfg, log),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.svr.RunServer(ctx)
}
