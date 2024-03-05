package main

import (
	"poulo-music/config"
	"poulo-music/logger"
	"poulo-music/server"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	log, err := logger.Init(cfg.Dir.BaseDir)
	if err != nil {
		panic(err)
	}

	app := server.NewApp(cfg, log)
	app.Run()
}
