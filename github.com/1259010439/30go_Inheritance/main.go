package main

import "fmt"

/*
  go语言中呢并没有真正意义上的继承，只能是进行模拟继承这种关系
  通过结构体的匿名内部结构体这种嵌套的形式来实现结构体之间的这种继承的关系
*/

//定义一个animal类
type animal struct {
	name string
	age int
}
//给anmial定义一个方法
func (a animal)work()  {
	fmt.Println(a.name,"在动！！！！！！！！！！！！")
}
//定义一个dog类继承animal
type dog struct {
	mouth int
	animal//通过将animal这个结构体写入dog的结构体进行一个继承的操作这样dog这个类就可以继承animal中的所有的方法和内部的参数
}
func (d dog)eat()  {
	fmt.Println(d.name,"吃啊")
}

func main()  {
 one:=dog{1,animal{"乐宝",2}}
 fmt.Println(one.age)//这个变量可以理解为从animal中继承过来的
 one.work()//这个方法就是从animal中继承过来的
 one.eat()//这个就是dog类自带的方法了
}
