package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/shyupc/go-study/crawler_distributed/config"
	"github.com/shyupc/go-study/crawler_distributed/rpcsupport"
	"github.com/shyupc/go-study/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	host := ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "https://album.zhenai.com/u/1106799723",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "雨宝",
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("error: %v", err)
	} else {
		fmt.Println(result)
	}
}
