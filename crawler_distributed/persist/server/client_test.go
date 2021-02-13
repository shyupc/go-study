package main

import (
	"testing"
	"time"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/model"
	"github.com/shyupc/go-study/crawler_distributed/config"
	"github.com/shyupc/go-study/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	host := ":1234"

	// Start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// Start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "https://album.zhenai.com/u/1106799723",
		Type: "zhenai",
		Id:   "1106799723",
		Payload: model.Profile{
			Name:       "雨宝",
			Gender:     "女",
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
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s;error: %v", result, err)
	}
}
