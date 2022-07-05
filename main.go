package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(
		//es服务不是跑在本地的，swtSniff=false, 不维护集群状态
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	//简单版本(近打印拉取信息)
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc: func(contents []byte) engine.ParseResult {
	//		return parser.ParseCityList(contents, 200) //爬取200个城市
	//	},
	//})

	//并发版
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount: 20,	//任务数
	//	ItemChan: persist.ItemSeaver(client),
	//}
	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc: func(contents []byte) engine.ParseResult {
	//		return parser.ParseCityList(contents, 200) //爬取200个城市
	//	},
	//})

	//队列版
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1,                          //任务数
		ItemChan:    persist.ItemSeaver(client), //处理响应结果的管道
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: func(contents []byte) engine.ParseResult {
			return parser.ParseCityList(contents, 470)
		},
	})
}
