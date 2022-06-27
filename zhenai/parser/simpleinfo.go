package parser

import (
	"crawler/engine"
	"crawler/model"
	"fmt"
	"regexp"
	"strconv"
)

const manProfile = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>男士</td> <td><span class="grayL">居住地：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([^<]+)</td> <!----> <td><span class="grayL">月   薪：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td> <td width="180"><span class="grayL">身   高：</span>([^<]+)</td></tr></tbody></table> <div class="introduce">([^<]+)</div>`
const womanProfile = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>女士</td> <td><span class="grayL">居住地：</span>([^<]+)</td></tr> <tr><td width="180"><span class="grayL">年龄：</span>([^<]+)</td> <td><span class="grayL">学   历：</span>([^<]+)</td> <!----></tr> <tr><td width="180"><span class="grayL">婚况：</span>([^<]+)</td> <td width="180"><span class="grayL">身   高：</span>([^<]+)</td></tr></tbody></table> <div class="introduce">([^<]+)</div>`

var (
	manMatches   = regexp.MustCompile(manProfile)
	womanMatches = regexp.MustCompile(womanProfile)
)

// ParseSimpleInfo 简单信息解析器
func ParseSimpleInfo(contents []byte) engine.ParseResult {
	profiles := make([]model.SimpleInfo, 0, 20)
	men := manMatches.FindAllSubmatch(contents, -1)
	for _, val := range men {
		age, _ := strconv.Atoi(string(val[4]))
		height, _ := strconv.Atoi(string(val[7]))
		profiles = append(profiles, model.SimpleInfo{
			Url:       string(val[1]),
			Nickname:  string(val[2]),
			Place:     string(val[3]),
			Age:       age,
			Income:    string(val[5]),
			Marriage:  string(val[6]),
			Height:    height,
			Introduce: string(val[8]),
		})
	}

	women := womanMatches.FindAllSubmatch(contents, -1)
	for _, val := range women {
		age, _ := strconv.Atoi(string(val[4]))
		height, _ := strconv.Atoi(string(val[7]))
		profiles = append(profiles, model.SimpleInfo{
			Url:           string(val[1]),
			Nickname:      string(val[2]),
			Place:         string(val[3]),
			Age:           age,
			EducationMate: string(val[5]),
			Marriage:      string(val[6]),
			Height:        height,
			Introduce:     string(val[8]),
		})
	}

	fmt.Println(profiles)
	return engine.ParseResult{}
}
