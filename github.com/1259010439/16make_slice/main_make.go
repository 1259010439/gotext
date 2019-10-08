package main

import "fmt"

//make 函数创建切片
func main()  {
	//	make([]T, size, cap) (切片类型，长度，容量)
	qp1 := make([]int,10,20)
	fmt.Printf("%v,%d,%d",qp1, len(qp1), cap(qp1))
	//判断一个切片是否为空应该用len()方法，不能用qp1==nil因为只要是切片初始化了就不会为空了
	fmt.Print(len(qp1)==0)
	b:=1
	for i,_:=range qp1{
		qp1[i]=b
		b++
	}
	fmt.Print(qp1)
}
