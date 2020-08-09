package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return nil,fmt.Errorf("wrong state code : %d",resp.StatusCode)
	}
	//包装bufio
	bodyReader := bufio.NewReader(resp.Body)
	//获取页面的编码格式
	//encoding := determineEncoding(bodyReader)
	//如果页面是gbk编码，需要进行转码
	//gbkReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//reader := transform.NewReader(resp.Body, encoding.NewDecoder())
	//return ioutil.ReadAll(reader)

	return ioutil.ReadAll(bodyReader)
}

//解析html中编码格式
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	peekBytes, err := r.Peek(1024)
	if err != nil{
		fmt.Printf("Fetcher error : %v",err)
		return unicode.UTF8
	}
	encoding, _, _ := charset.DetermineEncoding(peekBytes, "")
	return encoding
}
