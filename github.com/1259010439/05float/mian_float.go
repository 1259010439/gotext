package main

import "fmt"

func main()  {
	f1 :=1.23456
	fmt.Printf("%T",f1)//go中的默认的float类型是float64的类型
	f2 := float32(0.21)
	fmt.Printf("%T",f2)//想定义float32位的变量要自己定义
	//f1 = f2 并且float32 和 float64之间是不可以进行转换的


}
