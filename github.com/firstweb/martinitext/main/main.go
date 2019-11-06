package main

import "github.com/codegangsta/martini"

func main() {
	m := martini.Classic()          //创建一个martini的实例
	m.Get("/hello", func() string { //处理get请求 返回一个字符串
		return "hello "
	})
	m.Post("/login", func() string {

		return "login"
	})
	m.Run()
}
