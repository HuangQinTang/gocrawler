package engine

import (
	"log"
)

// QueueEngine 队列引擎
type QueueEngine struct {
	Scheduler   QueueScheduler //调度器
	WorkerCount int            //同时进行的任务数
}

// QueueScheduler 调度器
type QueueScheduler interface {
	Submit(request Request)                 //提交任务
	ConfigureMasterWorkerChan(chan Request) //工作的管道
	WorkerReady(chan Request)
	Run()
}

// Run 引擎启动
func (e *QueueEngine) Run(seeds ...Request) {
	//创建输出管道
	out := make(chan ParseResult)

	//启动队列任务，维护了2个队列，一个任务队列一个工作队列
	//任务会依次分发到工作队列里每个工人头上
	e.Scheduler.Run()

	//根据设置的任务数启动对应工作协程(创建workerCount个工作管道)
	for i := 0; i < e.WorkerCount; i++ {
		createQueueWorker(out, e.Scheduler)
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

// createQueueWorker 创建工作goroute
func createQueueWorker(out chan ParseResult, s QueueScheduler) {
	//创建当前协程的工作管道
	in := make(chan Request)
	go func() {
		for {
			//把当前工作管道放入调度器，等待分发任务
			s.WorkerReady(in)

			//任务分发到我(当前工作管道)头上时处理
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			//把处理结果传到输出管道
			out <- result
		}
	}()
}
