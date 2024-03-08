package config

import (
	"fmt"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	os.Setenv("poulo_mode", "debug")
	cfg, err := Init()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}

func TestMkdirCurrentPath(t *testing.T) {
	err := os.Mkdir("./dir", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMkdir(t *testing.T) {
	err := os.Mkdir("/Users/mulinbiao/workspace/poulo-music/dir/cache", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMkdirAll(t *testing.T) {
	err := os.MkdirAll("/Users/mulinbiao/workspace/poulo-music/dir/cache", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMkdirAllCurrentPath(t *testing.T) {
	err := os.MkdirAll("./dir", os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
}
