package parser

import (
	"fmt"
	"myGo/src/crawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	compile := regexp.MustCompile(cityListRe)
	//all := compile.FindAll(content, -1) //-1代表要所有的匹配
	submatch := compile.FindAllSubmatch(contents, -1)
	//for _, m := range all {
	//	fmt.Printf("%s\n", m)
	//}
	result := engine.ParseResult{}
	for _, m := range submatch {
		fmt.Printf("%s\n", m[1])
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
