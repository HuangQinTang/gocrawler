package client

import (
	"crawler/distributed/config"
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"crawler/engine"
)

func CreateProcessor(host string) (engine.Processor, error) {
	client, err := rpcsupport.NewClient(host)

	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		//入参序列化
		sReq := worker.SerializeRequest(req)

		//序列化后的对象作入参，调取对应rpc服务
		var sResult worker.ParseResult
		err = client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
