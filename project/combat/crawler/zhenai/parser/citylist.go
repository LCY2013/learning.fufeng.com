package parser

import (
	"learning.fufeng.com/project/combat/crawler/engine"
	"regexp"
)

//const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)" data-v-5e16505f[^>]*>([^<]+)</a>`
const cityListRe  = `<a data-v-[^>]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)">([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg:= regexp.MustCompile(cityListRe)
	submatch := reg.FindAllSubmatch(contents, -1)

	parseResult := engine.ParseResult{}
	for _,cont := range submatch {
		parseResult.Items = append(parseResult.Items, string(cont[2]))
		parseResult.Requests = append(parseResult.Requests,engine.Request{
			Url:       string(cont[1]),
			ParserFun: engine.NilParseFun,
		})
	}
	return parseResult
}
