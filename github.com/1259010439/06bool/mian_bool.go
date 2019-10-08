package main

import "fmt"

func main()  {
	b1 := true
	var b2 bool
	fmt.Print(b1);
	//go中的bool的默认值是false
	fmt.Printf("%v %T",b2,b2)
}