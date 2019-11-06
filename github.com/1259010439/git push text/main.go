package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://www.baidu.com"
	resq, err := http.Get(url)
	if err != nil {
		fmt.Println("错误发生")
	}
	fmt.Println("resq", resq)
}
