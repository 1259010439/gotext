package main

import "fmt"

func main()  {
	//普通的if
	age :=8
	if age>18 {
		fmt.Print("成年人")
	}else if age<=18 {
		fmt.Println("未成年人")
	}
	//go中特殊的if判断  p的作用域只在if中有效
	if p:=9;p>10 {
		fmt.Println("我是你爸爸")
	}else if p<10{
		fmt.Println("不我才是")
	}


}
