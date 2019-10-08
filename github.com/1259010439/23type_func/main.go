package main

import "fmt"

/*
go 语言中可以把函数做为另一个函数的参数，或者返回值
 */
func f1( m int) (x int)   {
	x = m
	return x
}
//参数是一个函数
func one(x func(int) int)  {
	fmt.Println(x(12))
}

//返回值是一个函数
func two(x int)(m func(int)int)  {
	return f1
}
func three(x int) (y int) {
	defer func() {
		x++
	}()
	defer func() {
		y=x
	}()
	return
}

func four() (int)  {
    x:=11
	defer func(x *int) {
		*x = 22
	}(&x)
	return x
}/*return的时候如果你没有给你的返回值进行命名的话，在底层会他会自己给函数的返回值进行一个命名
   你之后的所以给返回值赋值操作都可以理解为是给这个底层自动命名的返回值进行一个复制所以在你defer的时候
   你修改的只是你自己生成的这个x变量的值并不是底层的返回值的
   步骤：1.(返回值) =x=11 2.defer x=22 3.RET(返回值)=11 x=22
*/
func five() (x int)  {
	y:=11
	x=y
	defer func(m *int) {
		(*m)=22
	}(&x)
	return
}// 步骤：1.(返回值就是x) =11 2.defer (返回值x)=22 3.RET(返回值x)=22
func six() (x int)  {
	y:=11
	x=y
	defer func( x int) {
		x=22
	}(x)
	return
}//1.(返回值就是x) =11 2.defer (返回值x)=11 (在defer中的那个x)=22 3.RET(返回值x)=22
//这个地方他的defer中的x的值是和全局变量中的x命名是一样的但是在defer作用域中修改的就是这个下(defer中x)的值并不影响到全局中的x
//所以所x函数之间的参数的传递是值传递
func main()  {

//one(f1)
//x:=two(12)
//fmt.Println(x(12))
//fmt.Println(three(10))
fmt.Println(four())
fmt.Println(five())
	fmt.Println(six())
}
