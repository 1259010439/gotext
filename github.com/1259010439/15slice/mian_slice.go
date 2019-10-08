package main

import "fmt"

func main()  {
	//定义一个切片，切片和数组的区别呢就是切片的长度是可变的，
	var s1 []string
	fmt.Print(s1 == nil)
	//对切片初始化之后他的值就不为空了,就是相当于没有开辟内存空间
	var s2 = []string {"1","2","3"}
	fmt.Println(s2==nil)
	//查看长度和容量
	//这个切片的容量和长度都是由底层的数组来决定的，在定义一个切片的时候底层会先创建一个数组然后再包装称切片返回回来
	fmt.Println("s1 ",len(s1),cap(s1))
	fmt.Println("s2 ",len(s2),cap(s2))

	//2由数组得到切片
	arr := [...]int{1,2,3,4,5,6}
	//由数组切过来的这个切片的初始的容量是由所切数组的长度来和你所切的位置来决定的
	qp1 :=arr[0:5]//0包括但是5不包括 (左包不包)
	//这个切片的长度是5但是他的容量是6
	fmt.Printf("qp1 len(%d) cap(%d)",len(qp1),cap(qp1))
	fmt.Println(qp1)
	qp2 := arr[1:4]
	//这个切片的容量是3 但是他的通量是5  他的容量就是所切数组的长度（6）-切片切的位置（1）
	fmt.Printf("qp2 len(%d) cap(%d)",len(qp2),cap(qp2))
	fmt.Println(qp2)

	//var qp = []int {0,1,2,3,4,5,6,7}

	//切片再切片,当你的切片是从另一个切片切割而来的时候他的容量取决于
	// 你所切的那个切片的容量qp1（cap=6）-切片所切的位置(3) = 现在切割后的容量
	qp3 := qp1[3:5]
	fmt.Printf("qp3 len(%d) cap(%d)", len(qp3),cap(qp3))//cap = 3
	fmt.Println(qp3)

	//切片是引用类型，他的所以的数都是基于底层的数组来进行改变的
	arr1 :=[...]int{0,1,2,3,4,5,6,7,8,9}
	qp4 :=arr1[3:8]
    qp5 := qp4[:5]
    fmt.Println("arr1",arr1)
	fmt.Println("qp4",qp4)
	fmt.Println("qp5",qp5)
    //底层数组修改后切片的值也随之发生改变
    arr1[3] = 1300
	fmt.Println("garr1",arr1)
	fmt.Println("gqp4",qp4)
	fmt.Println("gqp5",qp5)
    //修改切片的值就是相当于修改底层数组的值
    qp5[0] = 20
	fmt.Println("garr1",arr1)
	fmt.Println("gqp4",qp4)
	fmt.Println("gqp5",qp5)

}
