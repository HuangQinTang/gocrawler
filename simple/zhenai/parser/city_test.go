package parser

import (
	"crawler/simple/fetcher"
	"fmt"
	"testing"
)

func TestParseCity(t *testing.T) {
	contentBts, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		t.Errorf("fetch error: %s", err.Error())
	}

	data := ParseCity(contentBts)
	for k, v := range data.Items {
		fmt.Println(v)
		fmt.Println(data.Requests[k].Url)
	}
}
