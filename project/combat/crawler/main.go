package main

import (
	"learning.fufeng.com/project/combat/crawler/engine"
	"learning.fufeng.com/project/combat/crawler/zhenai/parser"
)

func main()  {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParserFun: parser.ParseCityList,
	})
}

//func main() {
//	resp, err := http.Get("http://www.zhenai.com/zhenghun")
//	if err != nil{
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK{
//		fmt.Printf("Error status code %d \n",resp.StatusCode)
//	}
//	//获取页面的编码格式
//	//encoding := determineEncoding(resp.Body)
//	//如果页面是gbk编码，需要进行转码
//	//gbkReader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
//	//gbkReader := transform.NewReader(resp.Body, encoding.NewDecoder())
//	//readAll, err := ioutil.ReadAll(gbkReader)
//
//	readAll, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		panic(err)
//	}
//	//fmt.Printf("%s\n",readAll)
//
//	printCityList(readAll)
//}
//
////打印出所有的城市列表,城市列表解析器
//func printCityList(contexts []byte) {
//	reg:= regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-5e16505f[^>]*>([^<]+)</a>`)
//
//	/*
//	查询所有的a标签
//	contextsAll := reg.FindAll(contexts, -1)
//	for _,cont := range contextsAll{
//		fmt.Printf("%s\n",cont)
//	}
//	fmt.Printf("matches found : %d \n",len(contextsAll))
//	*/
//
//	submatch := reg.FindAllSubmatch(contexts, -1)
//	for _,cont := range submatch {
//		/*for _,subCont := range cont {
//			fmt.Printf("%s ",subCont)
//		}
//		fmt.Println()*/
//		fmt.Printf("City: %s, URL: %s \n",cont[2],cont[1])
//	}
//	fmt.Printf("matches found : %d \n",len(submatch))
//}
//
////解析html中编码格式
//func determineEncoding(r io.Reader) encoding.Encoding {
//	peekBytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil{
//		panic(err)
//	}
//	encoding, _, _ := charset.DetermineEncoding(peekBytes, "")
//	return encoding
//}