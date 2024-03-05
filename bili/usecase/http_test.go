package usecase

import (
	"fmt"
	"testing"
)

func Test_HTTP(t *testing.T) {
	http := NewHttp()
	http.SetUrl("https://www.bilibili.com")
	http.SetMethod("GET")

	http.AddParam("keyword", "hello word")

	fmt.Println(http.parameter.Encode())
}
