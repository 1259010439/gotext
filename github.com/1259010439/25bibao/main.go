package main

import "fmt"

/*
  闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
  就是一个内部函数 它可以调用他外部函数的变量这样就是闭包的原理
 */
//这就是一个典型的一个闭包的函数，f1的返回值是两个函数add 和 delete这两函数他们都有调用f1中的x参数
//这就是一个闭包
func f1(r int) ( add func(int)(int),delete func(int)(int))  {
    add = func(x int)(y int) {
    	r = r+x
    	return r
	}
    delete = func(x int)(y int) {
    	r =r-x
    	return r
	}
    return
}
func main()  {
    add,delete:=f1(10)//f1方法返回两个函数 add 和 delete
    fmt.Println(add(5),delete(5))//这两个函数又调用了f1中的x参数所以x的值会发生改变，
                                // add(5)=15 r = 10+5 r=15  delete(10)=10 r = 15-5 r= 10
    fmt.Println(add(4),delete(5))//add(4) r=10+4=14  delete(5) r = 14 - 5=9
	fmt.Println(add(3),delete(5))//...
}
