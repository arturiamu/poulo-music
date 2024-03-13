package kugou

import (
	"encoding/json"
	"fmt"
	"poulo-music/httpp"
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	var client = new(Kugou)

	data, err := client.search("爱一点", 1, 1)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)

	js, err := json.Marshal(data)
	fmt.Println(string(js))
}

func TestInfo(t *testing.T) {
	var client = new(Kugou)

	data, err := client.info("24518498c5d8fe15452b7d23532b0389")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}

func TestKugou(t *testing.T) {
	fmt.Println("vip \u65b0\u96c6\u7fa4(\u53ef\u591a\u6d3b)")
}

func TestKugouSearch(t *testing.T) {
	var url = "http://mobilecdn.kugou.com/api/v3/search/song?format=json&keyword=周杰伦&page=1&pagesize=20&showtype=1"
	bytes, err := httpp.NewHttp().SetUrl(url).Do()
	if err != nil {
		t.Error(err)
	}

	str := string(bytes)
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "/", "")

	fmt.Println(str)
}

func TestKugouNewSongs(t *testing.T) {
	var url = "http://m.kugou.com/?json=true"
	bytes, err := httpp.NewHttp().SetUrl(url).Do()
	if err != nil {
		t.Error(err)
	}

	str := string(bytes)
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "/", "")

	fmt.Println(str)
}

func TestKugouSongInfo(t *testing.T) {
	var url = "http://m.kugou.com/app/i/getSongInfo.php?cmd=playInfo&hash=24518498C5D8FE15452B7D23532B0389"
	bytes, err := httpp.NewHttp().SetUrl(url).Do()
	if err != nil {
		t.Error(err)
	}

	str := string(bytes)
	str = strings.ReplaceAll(str, "\r", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "/", "")

	fmt.Println(str)
}

func TestKugouSongLrc(t *testing.T) {
	var url = "http://m.kugou.com/app/i/krc.php?cmd=100&hash=24518498C5D8FE15452B7D23532B0389&timelength=3012000"
	bytes, err := httpp.NewHttp().SetUrl(url).Do()
	if err != nil {
		t.Error(err)
	}

	str := string(bytes)

	fmt.Println(str)
}
