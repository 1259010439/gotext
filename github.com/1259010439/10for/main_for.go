package main

import "fmt"

func main()  {
	//  for 初始语句;条件表达式;结束语句{
	//    循环体语句
	//  }
		for i:=1;i<10;i++{
			for j:=i;j<10;j++{
				fmt.Print(i,"*",j,"=",i*j," ")
			}
			fmt.Println(" ")
		}

		//for range(键值循环)
		s :="1231312321312我和我的祖国永远都不能分割"
		//这个k就是数组的小标但是一个中文会占用3个字节所以就很难受 想要取出中文就要先转换成runa数组在遍历
	   for key,value :=range s{
	   	fmt.Printf("%d+%c\n",key,value)
	   }
	   sum := []rune(s)
		for i:=0;i< len(sum);i++  {
			fmt.Printf("%d+%c ",i,sum[i])
		}
}
