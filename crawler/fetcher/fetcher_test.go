package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	fetch, err := Fetch("https://album.zhenai.com/u/1106799723")
	if err != nil {
		panic(err)
	}
	fmt.Println(fetch)
}
