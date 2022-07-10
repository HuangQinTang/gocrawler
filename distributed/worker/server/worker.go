package main

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"fmt"
	"log"
)

// 爬取服务
func main() {
	host := fmt.Sprintf(":%d", config.WorkerPort0)
	fmt.Println( "worker start"+host)
	log.Fatal(rpcsupport.ServerRpc(host, worker.CrawlService{}))
}
