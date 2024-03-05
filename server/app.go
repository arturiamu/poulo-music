package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"poulo-music/bili"
	biliUseCase "poulo-music/bili/usecase"
	"poulo-music/codec"
	codecUseCase "poulo-music/codec/usecase"
	"poulo-music/config"
	"poulo-music/musicify"
	"poulo-music/musicify/repository/localstorage"
	musicUseCase "poulo-music/musicify/usecase"
)

type App struct {
	biliUseCase  bili.UseCase
	musicUseCase musicify.UseCase
	codecUseCase codec.UseCase
}

func NewApp(cfg *config.Config, log *logrus.Logger) *App {
	db, err := initDB(cfg.Dir)
	if err != nil {
		panic(err)
	}

	musicRepo := localstorage.NewMusicLocalstorage(db)

	return &App{
		biliUseCase:  biliUseCase.NewUseCase(log, &cfg.Dir),
		musicUseCase: musicUseCase.NewUseCase(log, musicRepo),
		codecUseCase: codecUseCase.NewUseCase(log),
	}
}

func initDB(dir config.Dir) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s%s%s", dir.BaseDir, string(filepath.Separator), "data.db")), &gorm.Config{})
}

func (app *App) Run() {

}
