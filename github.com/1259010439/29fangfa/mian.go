package main

import "fmt"

type person struct {
	name string
	age int
}
//go 中的方法其实就是一个特殊的函数，该函数指定了具体调用它的类型只有该类型可以调用此方法
//其实这个方法就是用来丰富结构体的，struct可以理解为java中的一个类，方法就是对这个类进行一个丰富
//就是相当于写在java类中的一些方可以只供给这个结构体进行调用
func (p person)one()string  {
	s:= p.name+"学习"
	return s
}
func main()  {
	var first = person{
		name: "zhangjinzhen",
		age:  22,
	}
	fmt.Println(first.one())

}
