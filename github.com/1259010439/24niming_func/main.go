package main

import "fmt"

/*
  匿名函数
 */
//这就是一个匿名的函数，匿名函数大多数的时候都是再函数的内部来进行使用的因为
var f1 = func(i int,v int)int {
       return 0
}
func main()  {
	fmt.Println(f1(12,11))
	//大多时候的用法 ,函数在寻找变量的时候会先寻找自己的函数内部是否有定义，
	//然后再到全据变量中去寻找 所以当再次调用f1的时候就是返回的自己内部的这个定义出来的函数
	f1 := func(x int)(i int) {
		return 22
	}
	fmt.Printf("%d\n",f1(1))

	func(x int ){
		fmt.Println(x)
	}(1)

}
