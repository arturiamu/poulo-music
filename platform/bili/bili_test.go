package bili

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"poulo-music/config"
	"poulo-music/models"
	"testing"
)

func TestBili_GetSearch(t *testing.T) {
	client := NewBili(&logrus.Logger{}, &config.Config{
		BaseDir: "/Users/mulinbiao/workspace/poulo-music/dir",
	})
	searchTypeResult, err := client.GetSearchTypeResult(context.Background(), &models.SearchTypeParam{
		SearchType: "video",
		Keyword:    "周杰伦",
	}, false, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(searchTypeResult.Result)
}

func TestBili_GetHotContent(t *testing.T) {
	client := NewBili(&logrus.Logger{}, &config.Config{
		BaseDir: "/Users/mulinbiao/workspace/poulo-music/dir",
	})
	result, err := client.GetHotContent(context.Background(), models.GetHotContentParam{CachePic: true})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", result)
}

func TestBili_GetVideoRanking(t *testing.T) {
	client := NewBili(&logrus.Logger{}, &config.Config{})
	result, err := client.GetVideoRanking(context.Background(), false, "")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", result.List[0])
}
