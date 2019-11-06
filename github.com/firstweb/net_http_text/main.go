package main

import (
	"fmt"
	"net/http"
)

//利用给的net/http包搭建一个web服务器
func hello(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter 就是我要返回前端的值返回前端的东西 这就是两个结构体
	//*http.Request 这个就是前端穿过来的一个请求的具体的信息
	fmt.Println("ww", w.Header(), " ", w.WriteHeader)
	fmt.Println(r.RequestURI, " ", r.Response, " ", r.Body, " ", r.Method)
	//fmt.Fprintln(w,"ni de baba")
}
func main() {
	http.HandleFunc("/", hello)                       //注册一个路由访问的路径就是 /
	err := http.ListenAndServe("127.0.0.1:8000", nil) //开启监听这个端口
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

}
