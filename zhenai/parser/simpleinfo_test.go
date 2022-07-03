package parser

import (
	"crawler/model"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseSimpleInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("simpleinfo.txt")
	if err != nil {
		panic(err)
	}

	result := ParseSimpleInfo(contents)
	if len(result.Items) < 1 {
		t.Errorf("解析失败")
	}
	for _, v := range result.Items {
		info, ok := v.(model.SimpleInfo)
		if !ok {
			t.Errorf("解析失败")
		}
		if info.Id != 1396738538 {
			t.Errorf("数据解析不正确")
		}
		fmt.Printf("%+v", info)
		break
	}
}
