package parser

import (
	"crawler/simple/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>(.{2})</td>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1) //-1,匹配所有
	result := engine.ParseResult{}
	for _, m := range matches {
		//fmt.Println(string(m[0])) //完整
		//fmt.Println(string(m[1])) //url
		//fmt.Println(string(m[2])) //昵称
		//fmt.Println(string(m[3])) //性别
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NewParser,
		})
	}
	return result
}
