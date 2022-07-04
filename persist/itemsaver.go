package persist

import (
	"context"
	"crawler/model"
	"crawler/rediservice"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSeaver(client *elastic.Client) chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			userInfo, ok := item.(model.SimpleInfo)
			if !ok {
				log.Println("Item Saver err: 【item format error】")
				continue
			}

			//判断是否已保存过
			id := fmt.Sprintf("%d", userInfo.Id)
			unique, err := rediservice.RedisServer.ZhenaiDuplicate(id)
			if err != nil {
				log.Printf("Item Saver err: 【%v】", err)
				continue
			}
			if !unique {
				log.Printf("Item Saver info: 真爱用户【%v】已存在", id)
				continue
			}

			//保存数据至es
			if _, err := save(client, userInfo); err != nil {
				log.Printf("Item Save fail: 【%v】", err)
			}
			log.Printf("Item Saver: got item: #%d：【%v】", itemCount, item)
			itemCount++
		}
	}()
	return out
}

func save(client *elastic.Client, userInfo model.SimpleInfo) (docId string, err error) {
	resp, err := client.Index().Index("zhenai").BodyJson(userInfo).Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
