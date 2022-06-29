package main

import (
	"crawler/engine"
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

	//多任务版
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
	e := engine.QueueEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 20,	//任务数
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: func(contents []byte) engine.ParseResult {
			return parser.ParseCityList(contents, 200) //爬取200个城市
		},
	})
}
