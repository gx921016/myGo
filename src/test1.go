package main

import (
	"fmt"
	"runtime"
	"time"
)

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.baidu.com")
}
func main() {
	//var r Retiever
	//r = gaoxiang.Retriever{D: "aaaaa", Contents: "bbbbb"}
	//fmt.Println(download(r))
	var a [10]int
	fmt.Println("Running in", runtime.Version())
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
