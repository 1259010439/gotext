package main

import "fmt"

/*
  结构体
  这个结构体呢和java中的那给类呢差不多,但是结构体是值类型的，并不是引用类型的
 */
type person struct {
	name string
	age int
	sex string
	hobby []string
}

func f1(person person)  {
	person.name="王大壮"
}
func f2(person *person)  {
	(*person).name="王大壮"
}
func main()  {
	var one = person{name:"张小花",age:22,sex:"男" ,hobby: []string{"篮球","足球"}}
	fmt.Println(one)
	f1(one)//这个传给f1的其实是one的一个副本，其值的改变并不会影响到one
	fmt.Println(one)
	f2(&one)//这个传给f2的其实是one的指针，因为是指针所以根据指针修改one的内容就是和直接修改one是一样的
	fmt.Println(one)


	//匿名结构体
	//直接给这个结构体起一个名字，然后直接就可以使用这个结构体，这种都是定义在函数内部的，随函数的
	// 生成和销毁不会占用多余的内存空间
	var temp struct{
		x string
		y person
	}
	temp.x="1"
	temp.y = one
	fmt.Println(temp)

	
}
