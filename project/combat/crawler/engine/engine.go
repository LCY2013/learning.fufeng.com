package engine

import (
	"learning.fufeng.com/project/combat/crawler/fetcher"
	"log"
)

func Run(seeds ...Request)  {
	var requests []Request
	for _,v := range seeds{
		requests = append(requests,v)
	}

	for len(requests) > 0{
		//获取第一个slice元素
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching url : %s",r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetch: error " +
				"fetching url %s: %v",r.Url,err)
			continue
		}

		parseResult := r.ParserFun(body)
		//后面三个点表示将slice分个添加到slice里面
		requests = append(requests,parseResult.Requests...)

		//打印items
		for _,v := range parseResult.Items{
			log.Printf("Got item %v",v)
		}
	}
}
