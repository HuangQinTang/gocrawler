package engine

// Request 种子结构体
type Request struct {
	Url        string                   //目标url
	ParserFunc func([]byte) ParseResult //解析器方法(解析url所返回数据)
}

// ParseResult 解析器返回结构体
type ParseResult struct {
	Items    []interface{} //所解析的目标对象
	Requests []Request     //目标对象下还需要处理的种子
}

func NewParser([]byte) ParseResult {
	return ParseResult{}
}
