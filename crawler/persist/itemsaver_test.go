package persist

import (
	"encoding/json"
	"testing"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/model"

	"github.com/olivere/elastic"

	"golang.org/x/net/context"
)

func TestSave(t *testing.T) {
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

	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.100:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		panic(err)
	}
	const index = "dating_profile_test"
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", *resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("got %v; expected %v", actual, expected)
	}
}
