package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/unicode"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func Fetch(url string) ([]byte, error) {
	//获取网页信息
	host := strings.TrimSpace(url)
	resp, err := http.Get(host)
	if err != nil {
		panic(err)
	}
	//结束关闭Body
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrror status code: %d", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)
	//如果网站编码为GBK进行转码在读取
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetchers error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(
		bytes, "")
	return e
}
