package qq

import (
	"fmt"
	"poulo-music/models"
	"testing"
)

func TestSearch(t *testing.T) {
	var client = new(QQ)

	data, err := client.search("烟雨行舟", 10)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestQQ_GetMusic(t *testing.T) {
	var client = new(QQ)
	data, err := client.GetMusic(models.GetMusicParam{Identifier: "003953y22qBMUT"})
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}
