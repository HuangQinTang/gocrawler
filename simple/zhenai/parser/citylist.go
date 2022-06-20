package parser

import (
	"crawler/simple/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" data-v-1573aa7c>([^<]+)</a>`

// ParseCityList 城市列表解析器解析器
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1) //-1,匹配所有

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]), //url
			ParserFunc: ParseCity,
		})
		//fmt.Printf("URL: %s City: %s \n", m[1], m[2])
	}
	//fmt.Println(len(matches))
	return result
}
