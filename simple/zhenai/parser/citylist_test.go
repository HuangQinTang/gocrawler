package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist.txt")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	var expectedUrls = []string{"http://www.zhenai.com/zhenghun/nantou", "http://www.zhenai.com/zhenghun/changzhou", "http://www.zhenai.com/zhenghun/zhuzhou"}
	var expectedCities = []string{"南通","常州","株洲"}

	//验证城市列表是否470个 2022-06-10
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d Items; but had %d", resultSize, len(result.Items))
	}
	for _, v := range result.
}
