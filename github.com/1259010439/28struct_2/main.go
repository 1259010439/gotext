package main

import "fmt"

type person struct {
	name string
	age int
	sex string

}

//person的构造函数作用就是简化结构体的定义
//构造函数的名字一般都是由new开头的
//构造函数有两种，一种是直接返回一个新定义的结构体，另外一种是返回一个结构体的指针
//在结构体中的字段少的时候可以用直接返回结构体，
// 但是当结构体中的字段过多的时候就要返回这个结构体的指针因为这样会节省内存的开销
func newperson1(name string,age int,sex string) person {
	return person{
		name,
		age,
		sex,
	}

}
//返回构造体的指针
func newperson2(name string,age int,sex string) *person {
	return &person{
		name,
		age,
		sex,
	}

}

func f1(x *person)  {
	x.name = "王大壮"
}
func main()  {
	var x = new(person)//这个x是这个new出来的person变量的指针，相当于直接给这个对象直接开辟了内存空间
	//  并且在go语言中直接对指针类型操作的话就是相当于直接操作这个变量，这个地方就很类似java的变量赋值了
	x.name = "张二妮"
	f1(x)
	fmt.Print(x)

	//构造函数，go语言中struct结构体是没有自带的构造函数的所以可以自己写一个构造函数

}
