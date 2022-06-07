package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
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
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// determineEncoding 猜测r的编码类型
// 返回r的字符编码， 解析失败返回utf8
func determineEncoding(r io.Reader) encoding.Encoding {
	bts, err := bufio.NewReader(r).Peek(1024)
	if err != nil { //错误默认返回utf8
		log.Printf("determineEncoding Error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bts, "")
	return e
}
