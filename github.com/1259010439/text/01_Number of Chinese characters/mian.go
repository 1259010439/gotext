package main

import (
	"fmt"
	"unicode"
)

//判断字符串中的中文的数量
func main()  {
    s1 :="hello北京 你好lenovo"
    res := 0
    for _,v := range s1{
    	//主要就是这个unicode包的 Is方法就是判断字符是否和unicode中的字符相等的，中文的字符集就是unicode.Han
		if unicode.Is(unicode.Han,v) {
			res++
		}
	}
    fmt.Print(res)
}
