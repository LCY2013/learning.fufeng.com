package parser

import (
	"fmt"
	"io/ioutil"
	"learning.fufeng.com/project/combat/crawler/engine"
	"regexp"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err!=nil{
		panic(err)
	}

	//fmt.Printf("%s\n",contents)
	parseCityList := ParseCityList(contents)

	//定义结果集的slice
	results := engine.ParseResult{}
	maps := make(map[string]string)
	for i,v := range parseCityList.Items{
		if _, ok := maps[fmt.Sprintf("%v",v)]; !ok{
			results.Items = append(results.Items, fmt.Sprintf("%v",v))
			results.Requests = append(results.Requests, parseCityList.Requests[i])
		}
		maps[fmt.Sprintf("%v",v)] = fmt.Sprintf("%v",v)
	}

	const resultCityListSize = 470
	if len(results.Requests) != resultCityListSize{
		t.Errorf("result should have %d," +
			"but had %d",resultCityListSize,len(results.Requests))
	}

	if len(results.Items) != resultCityListSize{
		t.Errorf("result should have %d," +
			"but had %d",resultCityListSize,len(results.Items))
	}

}

//提取出@name 这种
const nameRe = "(@[^@]+)"
//const cityListRe  = `<a data-v-[^>]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)">([^<]+)</a>`
func TestRegexp(t *testing.T)  {
	reg := regexp.MustCompile(nameRe)
	submatch := reg.FindAllSubmatch([]byte("@阿斯达 @哈哈 @asdlsd@"), -1)
	for k,cont := range submatch {
		fmt.Println(k,"->",string(cont[1]))
	}
}
