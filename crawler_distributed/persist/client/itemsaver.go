package client

import (
	"log"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler_distributed/config"
	"github.com/shyupc/go-study/crawler_distributed/rpcsupport"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d %v", itemCount, item)
			itemCount++

			//Call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}