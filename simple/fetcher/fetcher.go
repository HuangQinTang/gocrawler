package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch 拉取网页url内容
// 返回[]byte数据
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetch Error: status code %s", resp.StatusCode)
	}

	//检测编码并替换为utf-8
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// determineEncoding 猜测r的编码类型
// 返回r的字符编码， 解析失败返回utf8
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bts, err := r.Peek(1024)
	if err != nil { //错误默认返回utf8
		log.Printf("determineEncoding Error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bts, "")
	return e
}
