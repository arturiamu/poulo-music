package config

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	cfg, err := Init()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
