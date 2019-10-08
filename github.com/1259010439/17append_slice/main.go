package main

import "fmt"

func main()  {
	//append这个方法是用来进行切片的增加，他的返回值就是一个切片，当切片的长度不够的时候他会有一个扩容的原则
	/*
		切片的扩容原则
		1.首先判断，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）。
		2.否则判断，如果旧切片的长度小于1024，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap），
		3.否则判断，如果旧切片长度大于等于1024，则最终容量（newcap）从旧容量（old.cap）开始循环增加原来的1/4，
	      即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		4.如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。
	*/
	s1 :=[]string{"哈尔滨","林口","北京"}
	s2 := append(s1,"南京")
	s1 =append(s1,"重庆")
	fmt.Println(s1,len(s1),cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Printf("%T",append(s1,"南京"))
	ss :=[]string{"牡丹江","济南","佳木斯","鹤岗"}
	//将切片添加到切片中去
	s1 = append(s1,ss...)//...是将切片进行一个拆分，能将切片中的所有元素都加到s1中去
	fmt.Println(s1,len(s1),cap(s1))

}
