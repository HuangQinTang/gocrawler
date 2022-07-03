package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSeaver(client *elastic.Client) chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item: #%d：%v", itemCount, item)
			itemCount++

			//保存数据至es
			if _, err := save(client, item); err != nil {
				log.Printf("Item Save fail: 【%v】", err)
			}
		}
	}()
	return out
}

func save(client *elastic.Client, item interface{}) (docId string, err error) {
	resp, err := client.Index().Index("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}

	return resp.Id, nil
}
