package parser

import (
	"regexp"

	"github.com/shyupc/go-study/crawler_distributed/config"

	"github.com/shyupc/go-study/crawler/engine"
)

var cityListRe = regexp.MustCompile(`<a href="(http://city.zhenai.com/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte, _ string) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
		//test one city
		break
	}
	return result
}
