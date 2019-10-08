package main

import "fmt"

func jump(x int)int  {
	if x==0{
		return 1
	} else if x==1 {
		return 1
	} else{
		return jump(x-1)+jump(x-2)
	}
}
func main()  {
	r:= jump(3)
	fmt.Print(r)
}
