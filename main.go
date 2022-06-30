package main

import (
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	//简单版本
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
		WorkerCount: 20, //任务数
		ItemChan:    persist.ItemSeaver(),	//处理响应结果的管道
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: func(contents []byte) engine.ParseResult {
			return parser.ParseCityList(contents, 1)
		},
	})
}
