package main

import (
	"context"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"poulo-music/bili"
	biliUseCase "poulo-music/bili/usecase"
	"poulo-music/codec"
	codecUseCase "poulo-music/codec/usecase"
	"poulo-music/config"
	"poulo-music/logger"
	"poulo-music/models"
	"poulo-music/musicify"
	"poulo-music/musicify/repository/localstorage"
	musicUseCase "poulo-music/musicify/usecase"
)

const (
	Sep = string(filepath.Separator)
)

// App struct
type App struct {
	biliUseCase  bili.UseCase
	musicUseCase musicify.UseCase
	codecUseCase codec.UseCase
	cfg          *config.Config
	ctx          context.Context
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

	db, err := initDB(cfg.Dir)
	if err != nil {
		panic(err)
	}

	musicRepo := localstorage.NewMusicLocalstorage(db)

	return &App{
		biliUseCase:  biliUseCase.NewUseCase(log, &cfg.Dir),
		musicUseCase: musicUseCase.NewUseCase(log, musicRepo),
		codecUseCase: codecUseCase.NewUseCase(log),
		cfg:          cfg,
	}
}

func initDB(dir config.Dir) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(fmt.Sprintf("%s%s%s", dir.BaseDir, string(filepath.Separator), "data.db")), &gorm.Config{})
}

func corsMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		h.ServeHTTP(w, r)
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	fs := http.FileServer(http.Dir(a.cfg.Dir.BaseDir + string(filepath.Separator) + "cache"))

	http.Handle("/", corsMiddleware(fs))

	http.ListenAndServe(":51730", nil)
}

func (a *App) GetBiliMusicRanking() (vo models.VO) {
	var cache = a.cfg.Dir.BaseDir + Sep + "cache" + Sep
	videoRankingData, err := a.biliUseCase.GetVideoRanking(a.ctx, true, cache)
	if err != nil {
		return vo.Fail(err.Error())
	}
	dataVO := videoRankingData.ToVO(a.cfg.App.FileServerAddr + Sep)
	return vo.Success(dataVO)
}

func (a *App) GetBiliSearch(keyword string, page int) (vo models.VO) {
	var cache = a.cfg.Dir.BaseDir + Sep + "cache" + Sep
	param := models.SearchTypeParam{SearchType: "video", Keyword: keyword, Page: page}
	searchTypeData, err := a.biliUseCase.GetSearchTypeResult(a.ctx, &param, true, cache)
	if err != nil {
		return vo.Fail(err.Error())
	}
	return vo.Success(searchTypeData.ToVO(a.cfg.App.FileServerAddr + Sep))
}

func (a *App) GetBiliAudio(bvid string) (vo models.VO) {
	var cache = a.cfg.Dir.BaseDir + Sep + "cache"
	video, err := a.biliUseCase.GetVideoViewByBvid(a.ctx, bvid)
	if err != nil {
		return vo.Fail(err.Error())
	}

	u, err := url.Parse(video.Pic)
	if err != nil {
		return vo.Fail(err.Error())
	}

	var music = models.Music{
		ID: 0, BVId: bvid, Name: video.Title, Title: video.Title,
		Pic: a.cfg.App.FileServerAddr + Sep + filepath.Base(u.Path),
	}

	var mp3 = cache + Sep + bvid + ".mp3"
	var mp4 = cache + Sep + bvid + ".mp4"

	_, err = os.Stat(mp3)
	if err != nil { //音频不存在
		_, err = os.Stat(mp4)
		if err != nil { //视频不存在
			err = a.biliUseCase.GetVideoMp4ByUri(a.ctx, bvid, mp4)
			if err != nil {
				return vo.Fail(err.Error())
			}
		}
		err = a.codecUseCase.GetAudioFromFile(mp4, mp3)
		if err != nil {
			return vo.Fail(err.Error())
		}
	}

	music.Url = a.cfg.App.FileServerAddr + Sep + bvid + ".mp3"
	return vo.Success(music)
}
