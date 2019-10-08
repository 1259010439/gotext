package main

import (
	"fmt"
)

func main()  {
	//找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	sum := [5]int{1, 3, 5, 7, 8}
	//for i,v:=range sum{
	//	for j,r:=range sum{
    //        if v+r==8{
	//			fmt.Println(i,j)
	//		}
	//	}
	//}

	for i:=0;i<5 ;i++  {
		for j:=i;j<5;j++{
			if sum[i]+sum[j]==8{
				fmt.Println(i,j)
			}
		}
	}

}
