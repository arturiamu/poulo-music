package qq

import (
	"context"
	"fmt"
	"os"
	"poulo-music/httpp"
	"poulo-music/models"
	"testing"
)

func TestSearch(t *testing.T) {
	var client = new(QQ)

	data, err := client.search("love story", 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestQQ_GetSearch(t *testing.T) {
	var client = new(QQ)

	// {ID:5105986 Platform:qq Identifier:001xd0HI0X9GNq Title:一路向北 Name:周杰伦 Describe:J III MP3 Player Type: Cover:https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg}
	data, err := client.GetSearch(context.Background(), models.GetSearchParam{
		Keyword:  "周杰伦",
		Pagesize: 100,
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestQQ_GetMusic(t *testing.T) {
	var client = new(QQ)

	// {ID:5105986 Platform:qq Identifier:001xd0HI0X9GNq Title:一路向北 Name:周杰伦 Describe:J III MP3 Player Type: Cover:https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg}
	data, err := client.GetMusic(context.Background(), models.GetMusicParam{Identifier: "001xd0HI0X9GNq"})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func Test_SaveLrc(t *testing.T) {
	bytes, err := httpp.NewHttp().SetUrl("http://www.ilingku.com/api/qqmusic/v2?act=lrcgc&cid=004OQ5Mt0EmEzv").Do()
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create("test.lrc")
	if err != nil {
		t.Error(err)
	}

	_, err = file.Write(bytes)
	if err != nil {
		t.Error(err)
	}
}
