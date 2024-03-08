package migu

import (
	"context"
	"fmt"
	"poulo-music/models"
	"testing"
)

func TestInfo(t *testing.T) {
	var client = new(Migu)

	data, err := client.info("60054704965")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestSearch(t *testing.T) {
	var client = new(Migu)

	data, err := client.search("周杰伦", 100)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestLrc(t *testing.T) {
	var client = new(Migu)

	data, err := client.lrcgc("60054701923")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestMigu_GetSearch(t *testing.T) {
	var client = new(Migu)

	//{ID:60054704965 Platform:qq Identifier:60054704965 Title:圣诞星（feat. 杨瑞代） Name:周杰伦 Describe: Type: Cover:https://d.musicapp.migu.cn/data/oss/resource/00/2t/1p/82511ebda9aa46289f583225286d943c}
	data, err := client.GetSearch(context.Background(), models.GetSearchParam{
		Keyword:  "周深",
		Pagesize: 100,
	})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestMigu_GetMusic(t *testing.T) {
	var client = new(Migu)

	// {ID:5105986 Platform:qq Identifier:001xd0HI0X9GNq Title:一路向北 Name:周杰伦 Describe:J III MP3 Player Type: Cover:https://y.gtimg.cn/music/photo_new/T002R300x300M000002MAeob3zLXwZ.jpg}
	data, err := client.GetMusic(context.Background(), models.GetMusicParam{Identifier: "69906401312"})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}
