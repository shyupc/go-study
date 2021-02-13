package parser

import (
	"regexp"
	"strconv"

	"github.com/shyupc/go-study/crawler/engine"
	"github.com/shyupc/go-study/crawler/model"
	"github.com/shyupc/go-study/crawler_distributed/config"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var genderRe = regexp.MustCompile("\"genderString\":\"([\u4e00-\u9fa5]+)\",\"hasIntroduce\"")
var heightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)kg</div>`)
var marriageRe = regexp.MustCompile("<div class=\"m-btn purple\" data-v-8b1eac0c>([\u4e00-\u9fa5]+婚)</div>")
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:(.*?)</div>`)
var xinzuoRe = regexp.MustCompile(`岁</div><div class="m-btn purple" data-v-8b1eac0c>(.*\(\d{2}\.\d{2}-\d{2}\.\d{2}\))</div>`)
var occupationRe = regexp.MustCompile("月收入:.*?</div><div class=\"m-btn purple\" data-v-8b1eac0c>([\u4e00-\u9fa5/-]+)</div>")
var educationRe = regexp.MustCompile("\"educationString\":\"([\u4e00-\u9fa5]+)\",\"emotionStatus\"")
var hukouRe = regexp.MustCompile("<div class=\"m-btn pink\" data-v-8b1eac0c>籍贯:([\u4e00-\u9fa5]+)</div>")
var carRe = regexp.MustCompile("<div class=\"m-btn pink\" data-v-8b1eac0c>([\u4e00-\u9fa5]+车)</div>")
var houseRe = regexp.MustCompile("<div class=\"m-btn pink\" data-v-8b1eac0c>([\u4e00-\u9fa5]+房)</div>")

var guessRe = regexp.MustCompile("<div class=\"m-btn pink\" data-v-8b1eac0c>([\u4e00-\u9fa5]+房)</div>")
var idUrlRe = regexp.MustCompile(`https://album.zhenai.com/u/([0-9]+)`)

func parseProfile(contents []byte, url, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = extractString(contents, genderRe)

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}

	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}

	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents, marriageRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Education = extractString(contents, educationRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)
	profile.Hukou = extractString(contents, hukouRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}

	matches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: NewProfileParser(string(m[2])),
		})
	}
	return result
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parser(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}
