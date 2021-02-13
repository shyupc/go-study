package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/shyupc/go-study/crawler_distributed/config"
	"github.com/shyupc/go-study/crawler_distributed/persist"
	"github.com/shyupc/go-study/crawler_distributed/rpcsupport"

	"github.com/olivere/elastic"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}

	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.10.100:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
