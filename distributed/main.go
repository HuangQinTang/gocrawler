package main

import (
	"crawler/distributed/config"
	itemServer "crawler/distributed/persist/client"
	worker "crawler/distributed/worker/client"
	"crawler/engine"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
)

func main() {
	//存储服务客户端
	itemChan, err := itemServer.ItemSeaver(fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	//爬取服务客户端
	processor, err := worker.CreateProcessor(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      1, //任务数
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}

	//启动调度器，调度器负责分发任务给客户端执行对应的服务
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: parser.NewParseCityList(10),
	})
}
