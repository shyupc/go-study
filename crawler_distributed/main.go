package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/scheduler"
	"github.com/shyupc/go-study/crawler/zhenai/parser"
	"github.com/shyupc/go-study/crawler_distributed/config"
	itemsaver "github.com/shyupc/go-study/crawler_distributed/persist/client"
	"github.com/shyupc/go-study/crawler_distributed/rpcsupport"
	worker "github.com/shyupc/go-study/crawler_distributed/worker/client"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemSaver, err := itemsaver.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemSaver,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    "http://city.zhenai.com/",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out
}
