package netease

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var client = new(Migu)

	data, err := client.search("周杰伦", 100)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", data)
}
