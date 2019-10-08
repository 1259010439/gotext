package main

import (
	"fmt"
	"strings"
)

func main()  {
	/*
	    查看这个字符串中单词出现的次数
	*/
	s1 := "what do you want to do"
	//先将字符串转换为切片
	qp1 := strings.Split(s1," ")
	fmt.Println(qp1)
	//定义一个map来 k是单词，v是单词的数量
	m := make(map[string]int)
	//遍历这个切片
	for  _,v:=range qp1  {
		//如果根据单词在map中取到了数量说明单词之前出现过，现在又出现了+1即可
		if _,ok:=m[v];ok{
             m[v]++
		}else{//如果没有出现过现在出现了那就给他一个值为1即可
            m[v] = 1
		}
	}

     for k,v:= range m{
     	fmt.Println(k,v)
	 }

}
