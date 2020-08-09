package parser

import (
	"learning.fufeng.com/project/combat/crawler/engine"
	"learning.fufeng.com/project/combat/crawler/model"
	"regexp"
	"strconv"
)

//const ageRex  = `[^<].* | ([\d]+)岁 | .* | .* | ([\d]+)cm | ([\d]+-[\d]+)元[^>]`
var ageRex = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var marriageRex = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}
	age, err := strconv.Atoi(extractString(contents,ageRex))
	if err == nil {
		profile.Age = age
	}
	profile.Marriage = extractString(contents,marriageRex)
	return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	submatch := re.FindSubmatch(contents)
	if len(submatch) > 2 {
		return string(submatch[1])
	}
	return ""
}
