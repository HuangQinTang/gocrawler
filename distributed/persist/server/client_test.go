package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":11001"
	//启动rpc服务
	go serverRpc(host, "zhenai")
	time.Sleep(time.Second)

	//创建rpc客户端
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	examples := model.SimpleInfo{
		Id:            888,
		Url:           "http://album.zhenai.com/u/1968078839",
		Nickname:      "山水有相逢",
		Gender:        "男士",
		Income:        "5001-8000元",
		Place:         "四川",
		Age:           18,
		EducationMate: "大专",
		Marriage:      "未婚",
		Height:        180,
		Introduce:     "本人是一个打工人生于1987年，普普通通的农民家庭",
	}

	var result string
	err = client.Call("ItemSaverService.Save", examples, &result)
	if err != nil {
		t.Error(err.Error())
	}
	if result == "" {
		t.Errorf("写入失败")
	}
	t.Log("执行成功，写入的文档id为" + result)
}
