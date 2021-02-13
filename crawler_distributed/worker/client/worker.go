package client

import (
	"net/rpc"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler_distributed/config"
	"github.com/shyupc/go-study/crawler_distributed/worker"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {
	return func(r engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(r)
		var sResult worker.ParseResult
		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}
