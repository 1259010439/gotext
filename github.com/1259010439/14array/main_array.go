package main

import "fmt"

//数组
	//存放元素的容器
	//和java一样必须指定长度和存放元素的种类
	func main()  {
		//GO的数组的长度也是数组的一部分

		//初始化也是和java一样的默认值都是一样的
		//1
		var sum = [3]int{1,2,3}
		fmt.Println(sum[0])
		//2根据初始值自动推断长度的大小
		 var sum1 = [...]int {1,2,3,4,5,56,1}
		fmt.Println(sum1[0])
		 //根据索引进行初始化
		 sum2:=[5]int{0:1,4:4}
		 fmt.Println(sum2[0],sum2[1],sum2[4])

		 //数组的遍历 和Java一样根据索引进行循环
		 home := [3]string{"林口","哈尔滨","北京"}
		for  i:=0;i< len(home);i++  {
               fmt.Println(home[i])
		}
		 //for range循环有点类似java中的foreach
		 for i,v := range home{
		 	fmt.Println(i,v)
		 }

		 //多维数组
		 arr := [3][4]int{{2,1},{2,4,3,5}}


		 for _,v:=range arr{
		 	fmt.Println(v)
		 	//for _,r:=range v{
		 	//	fmt.Println(r)
			//}
		 }

		 //go中的数组是值类型和java不一样java是引用类型所以这个数组
		 //每个数组都是相对独立的互不影响
		 b1:=[3]int{1,2,3}
		 b2:=b1
		 b2[0] = 100
		 fmt.Println("b1",b1)
		 fmt.Println("b2",b2)




	}
