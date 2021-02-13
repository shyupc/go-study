package main

import (
	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/persist"
	"github.com/shyupc/go-study/crawler/scheduler"
	"github.com/shyupc/go-study/crawler/zhenai/parser"
	"github.com/shyupc/go-study/crawler_distributed/config"
)

func main() {
	itemSaver, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
