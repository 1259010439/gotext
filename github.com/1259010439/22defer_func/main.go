package main

import "fmt"

/*
   defer
   函数中被这个关键字所修饰的语句会延迟到函数返回的时候执行
   如果有多个的话就按照添加defer关键字的倒叙执行 类似于一个队列
   在数据库链接释放的时候java也有类似的操作，用这个关键字就可以做到。
 */
func one()  {
	fmt.Println("开始")
	fmt.Println("wo cao")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("结束")
} /*
    go中的return的操作不是原子性的 他是由两步骤组成的
    1.返回值赋值
    2.在执行ret（就是跳出循环的操作）操作
    如果在函数中存在defer关键字
    1.返回值赋值
    2.执行defer
    3.在执行ret（就是跳出循环的操作）操作
  */
func two()int  {
	x:=12//x的值已经被附上了所以在return的时候就是返回的这个x的值
	defer func() {
		x=22//这个操作只是修改了x的值并不是修改了返回要用到的值
	}()
	return x
}

func three() (x int)  {
	//这个返回值就是25 因为在函数定义的时候返回值就和x进行了一个绑定 返回值就是x所以修改了x的值就是修改了返回值
	defer func() {
		x=25
	}()
	return 12
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main()  {
	one()
	fmt.Println(two())
	fmt.Println(three())
	s:=three//这个样子不带（）就得到的就是这个函数
	fmt.Printf("%d %T\n",s,s)
	fmt.Println("==================================================")
    //defer的面试题 defer在修饰的方法和语句在进栈的时候会按照代码的顺序进行入栈，并且变量的赋值也跟着一起进入栈中的

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	/*入栈的顺序
	  1.a=1
	  2.b=2
	  3.calc("10",1,2) = 3 输出 10 1 2 3 因为这个第一个calc中的calc是一个变量所以要先给defer修饰的这个变量进行赋值
	    所以这个语句就会先执行
	  4. calc("1",1,(3.的返回值3)) = 4 输出1 1 3 4  这个就会压进栈中
	  5. a=0
	  6.calc("20",0,2) = 2 输出20 0 2 2同理这个也要先执行
	  7.calc("2",0,(6.的返回值2)) = 2 输出2 0 2 2
	  8.b=1

	 //所以输出的顺序就是
	 3，输出  10 1 2 3
	 6，输出  20 0 2 2
	 然后就是按照栈进行输出 进栈顺序是4 ， 7
	输出顺序就是 7， 4
	7.2 0 2 2
	4.1 1 3 4

	 */

}
