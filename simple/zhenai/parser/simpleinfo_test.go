package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseSimpleInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("simpleinfo.txt")
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		panic(err)
	}

	ParseSimpleInfo(contents)
}
