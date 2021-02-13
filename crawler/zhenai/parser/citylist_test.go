package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(content, "")

	const resultSize = 470
	expectUrls := []string{
		"http://city.zhenai.com/aba",
		"http://city.zhenai.com/akesu",
		"http://city.zhenai.com/alashanmeng",
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
