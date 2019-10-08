package main

import "fmt"

const (
	n1 = iota//0
	n2       //1
	n3       //2
)

const (
	//iota在常量出现的时候会先赋一个初值0，随后每次增加一行就加1
	 d1 = iota//0 iota = 0
	 d2 = 100//100 iota =1
	 d3      //100  iota=2
	 d4 = iota//3   iota = 3
)
const(
	s1,s2,s3 = iota,iota+1,iota //这一行iota值为0 所以 s1=0,s2=1,s3=1
	s4,s5,s6 = iota,iota,iota//这一行 iota的值为1 所以 s4=1,s5=1,s6=1

)

func main()  {
	fmt.Println("n",n1,n2,n3)
	fmt.Println("d",d1,d2,d3,d4)
	fmt.Println("s",s1,s2,s3,s4,s5,s6)
}
