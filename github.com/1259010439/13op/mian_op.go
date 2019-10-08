package main

import "fmt"

func main()  {
	var(
		a=10
		b=3
	)
	//算数运算符
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	a++//在go中是单独执行的 不能像java中的 b=a++可以将a++的值赋值给别的变量
	a--//和a++一样都是这种操作

	//关系运算符
	//这些运算符的用法和java中都一样
	fmt.Println(a == b)//Go是强类型的语言和java一样所以只有同类型的才能进行比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a < b)

	//逻辑运算符
	age:=22
	//并且
	if age>18 && age<60{
		fmt.Println("还没退休")
	}else{
		fmt.Println("不用上班了")
	}
    //或者
	if age>18||age<60{
		fmt.Println("上班把！！")
	}else {
		fmt.Println("不用上班了")
	}
    //取反
    falg:=true
    fmt.Println(!falg)

    //位运算符
    //&与运算 两位均为1才是1
    /*
    *5： 101
     2： 010
     &   000
    */
    fmt.Println("5&2",5&2)

	//|按位或运算 两位有一个1就是1
	/*
	   *5： 101
	    2： 010
	    |   111
	*/
	fmt.Println("5|2",5|2)

	//^按位 异或运算 两位不一样就是1
	/*
	   *5： 101
	    2： 010
	    |   111
	*/
	fmt.Println("5^2",5^2)
	// <<将二进制的数向左移动 相当于*2的次方
	/*
	  5： 101
	 <<1 1010 10
	 <<2 10100 20
	 */
	fmt.Println("5<<1",5<<1)
	fmt.Println("5<<2",5<<2)
	//>>将二进制的数据向右移动 相当于/2的次方
	/*
	  5： 101
	 >>1  10 2
	 >>2   1 1
	*/
	fmt.Println("5>>1",5>>1)
	fmt.Println("5>>2",5>>2)


	
}
