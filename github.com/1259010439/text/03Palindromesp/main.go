package main

import "fmt"

func main()  {
	/*
	  回文数的判断
	  上海自来水来自海上 就是回文数
	*/
	var s = "123321"
	for i:=0;i<len(s)/2;i++ {
		if s[len(s) - i - 1] !=s[i]{
			fmt.Println("我不是回文")
			return
		}
		fmt.Print("我是回文")
	}
}
