package engine

import "crawler/model"

// Request 种子结构体
type Request struct {
	Url        string //目标url
	Parser     Parser //解析器方法(解析url所返回数据)
}

// ParseResult 解析器返回结构体
type ParseResult struct {
	Items    []model.SimpleInfo //所解析的目标对象(这里是简单用户信息）
	Requests []Request          //目标对象下还需要处理的种子
}

type ParserFun func(contents []byte) ParseResult

type Parser interface {
	Parse(contents []byte) ParseResult
	Serialize() (name string, args interface{})
}

type NilParser struct {
}

func (NilParser) Parse(_ []byte) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFun
	name   string
}

func (f *FuncParser) Parse(contents []byte) ParseResult {
	return f.parser(contents)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

func NewFuncParser(p ParserFun, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
