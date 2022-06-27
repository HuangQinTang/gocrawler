package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// Submit 把任务提交到工作的管道
func (s *SimpleScheduler) Submit(r engine.Request) {
	//避免循环阻塞，这里用协程
	go func() { s.workerChan <- r }()
}

// ConfigureMasterWorkerChan 配置工作管道
func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
