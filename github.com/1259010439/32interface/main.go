package main

import "fmt"

/*
  interface 接口  **接口其实就是一种类型**
  如果一个一个类型（结构体，可以理解为Java中的一个类）它实现了接口中的所有的方法
  那么他就是实现了这个接口，他就拥有了和这个接口一样的数据类型，
  实现接口就是相当于一种拓展，让该类型同时拥有多种的数据类型，
  并且在功能上和继承的接口的数据类型完全一样

  接口类型的变量的具体类型是和值一样需要进行赋值的，在初始化的时候他的类型其实是为<nil>的
  因为接口变量他的类型和值都是动态的，所以接口类型的变量的存储值的方式是分为两部分的
  第一部分：存放所传进来变量的类型
  第二部分：存放所传进来的值
 */
type animal interface {
	eat()
}

//dog这个类型它实现了eat的这个方法所以它现在可以看做时候一个animal类型
type dog struct {
	name string
}
func (d dog)eat()  {
	fmt.Println(d.name+"狗在吃！！！")
}
//Cat
type cat struct {
	name string
}
func (c cat)eat()  {
	fmt.Println(c.name+"猫吃啊")
}

func wei(a animal)  {
	a.eat()
}

func main()  {
	var one dog
	fmt.Printf("%T\n",one)//main.dog 普通类型在初始化类型就已经确定了
	one=dog{name:"乐宝"}
	two:=cat{"喵喵"}
    wei(one)
	wei(two)
	var a1 animal
	fmt.Printf("%T\n",a1)//<nil>因为这个时候变量刚刚初始化并没有赋值，接口类型的变量是需要进行赋值的所有这个时候的变量类型就是nil
	a1 = one
	fmt.Printf("%T\n",a1)//main.dog这个时候给接口类型进行了一个赋值所以他的类型也就进行了赋值
	a1=two
	fmt.Printf("%T\n",a1)//mian.cat对a1再一次进行赋值他的类型也发生了改变

}