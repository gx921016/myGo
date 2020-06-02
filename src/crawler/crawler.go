package main

import (
	"myGo/src/crawler/engine"
	"myGo/src/crawler/zhenai/parser"
)

func main() {
	//获取网页信息
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
