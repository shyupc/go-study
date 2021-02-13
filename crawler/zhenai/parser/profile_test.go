package parser

import (
	"io/ioutil"
	"testing"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")

	if err != nil {
		panic(err)
	}

	result := parseProfile(contents, "https://album.zhenai.com/u/1106799723", "雨宝")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v ", result.Items)
	}

	actual := result.Items[0]
	expected := engine.Item{
		Url:  "https://album.zhenai.com/u/1106799723",
		Type: "zhenai",
		Id:   "1106799723",
		Payload: model.Profile{
			Name:       "雨宝",
			Gender:     "女士",
			Age:        28,
			Height:     163,
			Weight:     51,
			Income:     "3-5千",
			Marriage:   "未婚",
			Education:  "高中及以下",
			Occupation: "自由职业",
			Hukou:      "江苏连云港",
			Xinzuo:     "巨蟹座(06.22-07.22)",
			House:      "租房",
			Car:        "已买车",
		},
	}

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
