package main

import (
	"fmt"
)

func main()  {
	s1 :=[]int{1,2,3}
	s2 :=s1
	//一定要初始化
	var s3 = make([]int,3,3)
	//copy 的使用方法 copy(要放到的切片，要赋值的切片)
	copy(s3,s2)
   fmt.Println(s1,s2,s3)

	//copy方法相当于复制一个这个值，把你这个切片中的数取出来然后再放放到一个新的切片中去
	//所有修改原切片的值不会影响已经复制好的切片
	s1[0]=12341
	fmt.Println(s1,s2,s3)
	//切片元素的删除
	qp := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(qp)
	//这个就是切片的删除
	qp = append(qp[:3],qp[4:]...)
	fmt.Print(qp)

	//切片的删除其本质上是对他的底层的数组进行了一个数值的调整，
	// 因为切片就是一个指针也就是一个“框”它里面显示出的数都是框中的数
	arr :=[...]int{1,2,3,4,5,6,7,8,9}
    a := arr[:]
    a = append(a[:2],a[7:]...)
    fmt.Println(arr,a)

}
