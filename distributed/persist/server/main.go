package main

import (
	"crawler/distributed/config"
	"crawler/distributed/persist"
	"crawler/distributed/rpcsupport"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func main() {
	log.Fatal(serverRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.Zhenai))
}

// serverRpc 启动itemsaver服务
// host 服务端口 esIndex 数据存储至es的索引名
func serverRpc(host string, esIndex string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return rpcsupport.ServerRpc(host, &persist.ItemSaverService{Client: client, EsIndex: esIndex})
}
