package engine

import (
	"log"
)

// ConcurrentEngine 多任务引擎
type ConcurrentEngine struct {
	Scheduler   Scheduler //调度器
	WorkerCount int       //同时进行的任务数
}

// Scheduler 调度器
type Scheduler interface {
	Submit(request Request)                 //提交任务
	ConfigureMasterWorkerChan(chan Request) //工作的管道
}

// Run 引擎启动
func (e *ConcurrentEngine) Run(seeds ...Request) {
	//创建输入、输出、任务管道
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in) //配置工作任务管道

	//监听任务管道开始干活
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//把要干的活提交到任务管道
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//输出结果
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Requests {
			log.Printf("Got item: #%d：%v", itemCount, item)
			itemCount++
		}

		//如果还有要做的任务继续提交
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// createWorker 创建工作goroute
func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
