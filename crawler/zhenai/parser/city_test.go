package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	content, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCity(content, "")

	const resultSize = 20
	expectUrls := []string{
		"http://album.zhenai.com/u/1106799723",
		"http://album.zhenai.com/u/1743408180",
		"http://album.zhenai.com/u/1817535558",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}
