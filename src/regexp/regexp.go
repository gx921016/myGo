package main

import (
	"fmt"
	"regexp"
)

const text = "assdfsfsf呆逼马sfsfsfsf"

//^[一-龥]{2,5}&
func main() {
	re := regexp.MustCompile("[一-龥]{2,8}")
	match := re.FindString(text)
	fmt.Println(match)
}
