package localstorage

import (
	"context"
	"crypto/sha256"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestHash(t *testing.T) {
	hash := sha256.Sum256([]byte("data"))
	fmt.Printf("%x\n", hash)
}

func TestLocalstorage(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	repo := NewMusicLocalstorage(db)
	err = repo.AutoMigrate()
	if err != nil {
		panic(err)
	}

	//music := &models.Music{BVId: "BV1bK4y1C7tL", Name: "test music", TimeSkips: []models.TimeSkip{{Timestamp: 1000, Time: 500}}}

	//music, err = repo.SaveMusic(music)
	//if err != nil {
	//	panic(err)
	//}

	//music1, err := repo.GetMusicByID(1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v\n", music1)

	//item, err := repo.SavePlaylistItem(&models.PlaylistItem{PlaylistID: 1, MusicID: 1})
	//fmt.Printf("%+v\n", item)

	musics, err := repo.ListMusicByPlaylistID(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	for _, music := range musics {
		fmt.Printf("%+v\n", music)
	}
}
