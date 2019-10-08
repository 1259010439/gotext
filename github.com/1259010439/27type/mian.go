package main

import "fmt"

//go语言支持自定义一个类型，但是这个类型只能再本包下进行使用
type String string
//类型别名，就是给go语言基本数据类型起一个其他的名字，就是小名字的意思其本质上还是原来的这个类型
type Integer = int
func main()  {
	var name String
	name = "gushihao"
	fmt.Printf("%T %v\n",name,name)//它的类型就是mian.String了

	var age Integer
	age = 22
	fmt.Printf("%T %v\n",age,age)//它的类型还是int

}
