package main

import "fmt"
//函数和java一样几乎

//带返回值和带参数的函数
func sum(a int,b int)int  {
	return  a+b
}

/*
无参数切有多个返回值的函数,在go中是可以给返回值取名字的 有了名字的返回值可以在程序中直接return
go中的函数还可以有多个返回值，要用，进行隔开 返回的时候return也要写多个
*/
//给返回值命名的函数
func f1()(res1 int,res2 string)  {
    res1 = 1
    res2 ="就是我1"
	return
}
//为给返回值命名的函数
func f2()( int, string)  {
	res3 := 2
	res4 :="就是我2"
	return res3,res4
}


func main()  {
	var a int
	var b int
	a = 11
	b = 12
	fmt.Println("sum:",sum(a,b))
	res1, res2 := f1()
	fmt.Println("f1:", res1,res2)
	res3,res4 := f2()
	fmt.Println("f2:", res3,res4)
}