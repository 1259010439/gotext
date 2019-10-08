package main

import "fmt"

func main()  {
	//整型
	var a = 101
	fmt.Printf("十%d\n",a)//十进制
	fmt.Printf("二%b\n",a)//二进制
	fmt.Printf("八%o\n",a)//八进制
	fmt.Printf("十六%x\n",a)//十六进制
	//八进制
	a2:=077
	fmt.Printf("十%d\n",a2)//十进制
	fmt.Printf("二%b\n",a2)//二进制
	fmt.Printf("八%o\n",a2)//八进制
	fmt.Printf("十六%x\n",a2)//十六进制
	//定义十六进制
	a3:= 0x1234f
	fmt.Printf("十%d\n",a3)//十进制
	fmt.Printf("二%b\n",a3)//二进制
	fmt.Printf("八%o\n",a3)//八进制
	fmt.Printf("十六%x\n",a3)//十六进制
	fmt.Printf("类型%T\n",a3)//变量的类型

	a4:=int8(18)
	fmt.Printf("%T%d\n",a4,a4)


}