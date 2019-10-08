package main

import "fmt"

func main()  {
	a := 22
	b :="123"
	c :="123"
	fmt.Print(b==c)
   //switch case的用法	go的switch中是可以写判断的但是只要有一个跳出就break了
   //fallthrough 不管下一个case的条件是否成立都会执行下一个case
	switch {
	case a<10:
		fmt.Println("我大于10")
	case  a<15:
		fmt.Print("我大于15")
	case a>20:
		fmt.Println("我大于20")
		fallthrough
	default:
		fmt.Print("我啥也不是")
	}
}
