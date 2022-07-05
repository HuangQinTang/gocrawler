package main

import (
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"crawler/model"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	log.Fatal(serverRpc(":10001", model.Zhenai))
}

func serverRpc(host string, esIndex string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{Client: client, Index: esIndex})
}
