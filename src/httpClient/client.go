package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/",
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "aaaaa7777你是个呆逼 没鸡鸡aaaaa")
			//path := request.URL.Path[len("/")]
			//file, err := os.Open(path)
			//if err != nil {
			//	panic(err)
			//}
			//defer file.Close()

		})
	http.ListenAndServe(":8888", nil)
}
