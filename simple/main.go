package main

import (
	"crawler/simple/engine"
	"crawler/simple/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc: func(contents []byte) engine.ParseResult {
			return parser.ParseCityList(contents, 200) //爬取200个城市
		},
	})
}
