package parser

import (
	"regexp"
	"strings"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler_distributed/config"
)

var cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*><span>([^<]+)</span></a>`)
var cityUrlRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*><span>([^<]+)</span></a>`)

func ParseCity(contents []byte, _ string) engine.ParseResult {
	matches := cityRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    strings.ReplaceAll(string(m[2]), "http://", "https://"),
			Parser: NewProfileParser(string(m[2])),
		})
	}

	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		//result.Items = append(result.Items, "User "+m[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, config.ParseCity),
		})
	}
	return result
}
