package main

import "fmt"

func main()  {
	/*
	 map类型
	 */
	//map 和 slice一样都是引用数据类型，所以在使用之前都要进行一个初始化的操作 用make或者直接赋值
	var m1   =make(map[int]string,10)
	m1[0] = "黑龙江"
	m1[1] = "牡丹江"
	m1[2] = "林口"
    //map的遍历 可以用这个range的遍历来实现
    for k,v:=range m1{
    	fmt.Println(k,v)
	}
    //只遍历value
	for _,m:=range m1{
		fmt.Println(m)
	}
	//只遍历key
	for s,_:=range m1{
		fmt.Println(s)
	}
	fmt.Print(m1)
	//删除map中的一个元素 用delete方法直接指定key删除就行
	delete(m1,1)
	fmt.Println(m1)
	//切片和map的组合 就是相当于java中的数组里面套map
	//值为map类型的切片
	var qp = make([]map[int]string,10,20)
	qp[0] = make(map[int]string,3)
	qp[0][11] = "长江"
	fmt.Println(qp)
    //值为切片的map
	var map1  = make(map[int][]string)
	map1[22] = []string{"1","3","4"}
	fmt.Print(map1)


}
