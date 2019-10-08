package main

import "fmt"

func main()  {
	//这个是和Java中有区别的break剩下的都一样
	// 这个是和Java中有区别的break剩下的都一样

	// ONE:
	//		for i:=0;i<10 ;i++  {
	//			fmt.Print(i)
	//			if i==5{
	//				break ONE
	//			}
	//
	//		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if j == 2 {
					// 设置退出标签
					goto breakTag
				}
				fmt.Printf("%v-%v\n", i, j)
			}
		}
		return
		// 标签
	breakTag:
		fmt.Println("结束for循环")

}
