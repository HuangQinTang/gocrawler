package parser

import (
	"crawler/simple/engine"
	"crawler/simple/model"
	"fmt"
	"regexp"
	"strconv"
)

const (
	baseInfoRe = `<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`
	livingRe   = `<div data-v-8b1eac0c="" class="m-btn pink">([^<]+)</div>`
	foodRe     = `<div data-v-8b1eac0c="" class="question f-fl">喜欢的一道菜：</div> <div data-v-8b1eac0c="" class="answer f-fl">([^<]+)</div>`
	idolRe     = `<div data-v-8b1eac0c="" class="question f-fl">欣赏的一个名人：</div> <div data-v-8b1eac0c="" class="answer f-fl">([^<]+)</div>`
	songRe     = `<div data-v-8b1eac0c="" class="question f-fl">喜欢的一首歌：</div> <div data-v-8b1eac0c="" class="answer f-fl">([^<]+)</div>`
	bookRe     = `<div data-v-8b1eac0c="" class="question f-fl">喜欢的一本书：</div> <div data-v-8b1eac0c="" class="answer f-fl">([^<]+)</div>`
	thingRe    = `<div data-v-8b1eac0c="" class="question f-fl">喜欢做的事：</div> <div data-v-8b1eac0c="" class="answer f-fl">([^<]+)</div>`
)

var (
	baseInfoMatches = regexp.MustCompile(baseInfoRe)
	livingMatches   = regexp.MustCompile(livingRe)
	foodMatch       = regexp.MustCompile(foodRe)
	idolMatch       = regexp.MustCompile(idolRe)
	songMatch       = regexp.MustCompile(songRe)
	bookMatch       = regexp.MustCompile(bookRe)
	thingMatch      = regexp.MustCompile(thingRe)
)

const (
	Marriage = iota
	Age
	Xinzuo
	Height
	Weight
	Workplace
	Income
	Occupation
	Education

	Nationality = iota
	Hokou
	Stature
	Smoking
	Drink
	House
	Car
	Child
	likeChild
	MarriageTime
)

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}

	//基本信息赋值
	matches := baseInfoMatches.FindAllSubmatch(contents, -1) //-1,匹配所有
	for k, val := range matches {
		switch k {
		case Marriage:
			profile.Marriage = string(val[1])
		case Age:
			ageBts := val[1][0 : len(val[1])-3] //去除末尾中文"岁"
			age, err := strconv.Atoi(string(ageBts))
			if err == nil {
				profile.Age = age
			}
		case Xinzuo:
			profile.Xinzuo = string(val[1])
		case Height:
			heightBts := val[1][0 : len(val[1])-2] //去除末尾"cm"
			height, err := strconv.Atoi(string(heightBts))
			if err == nil {
				profile.Height = height
			}
		case Weight:
			weightBts := val[1][0 : len(val[1])-2] //去除末尾"kg"
			weight, err := strconv.Atoi(string(weightBts))
			if err == nil {
				profile.Weight = weight
			}
		case Workplace:
			workplace := val[1][10:] //去除开头"工作地:"
			profile.Workplace = string(workplace)
		case Income:
			income := val[1][10:] //去除开头"月收入:"
			profile.Income = string(income)
		case Occupation:
			profile.Occupation = string(val[1])
		case Education:
			profile.Education = string(val[1])
		}
	}

	//生活状况
	matches = livingMatches.FindAllSubmatch(contents, -1)
	for k, val := range matches {
		switch k {
		case Nationality:
			profile.Nationality = string(val[1])
		case Hokou:
			hokou := val[1][7:] //去除开头的"籍贯"
			profile.Hokou = string(hokou)
		case Stature:
			stature := val[1][7:] //去除开头"体型"
			profile.Stature = string(stature)
		case Smoking:
			profile.Smoking = string(val[1])
		case Drink:
			profile.Drink = string(val[1])
		case House:
			profile.House = string(val[1])
		case Car:
			profile.Car = string(val[1])
		case Child:
			profile.Child = string(val[1])
		case likeChild:
			likeChildBts := val[1][19:] //去除开头"是否想要孩子:"
			profile.LikeChild = string(likeChildBts)
		case MarriageTime:
			marriageTime := val[1][13:]
			profile.MarriageTime = string(marriageTime)
		}
	}

	//喜欢的食物
	food := foodMatch.FindSubmatch(contents)
	profile.Food = string(food[1])

	//喜欢的名人
	idol := idolMatch.FindSubmatch(contents)
	profile.Idol = string(idol[1])

	//喜欢的歌
	song := songMatch.FindSubmatch(contents)
	profile.Song = string(song[1])

	//喜欢的书
	book := bookMatch.FindSubmatch(contents)
	profile.Book = string(book[1])

	//喜欢做的事
	thing := thingMatch.FindSubmatch(contents)
	profile.Hobby = string(thing[1])

	fmt.Printf("%#v", profile)
	return engine.ParseResult{}
}
