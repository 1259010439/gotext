package main

import "fmt"
//Println :可以打印出字符串，和变量
//Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量
//总结来说就是printf要在需要格式化的地方输出，其他直接输出变量或者直接输出字符串的可以直接用println

//fmt的占位符
func main()  {
	a3 := 100
	//查看类型
	fmt.Printf("%T\n",a3)
	//查看值
	fmt.Printf("%v\n",a3)
	//十进制
	fmt.Printf("十%d\n",a3)
	//二进制
	fmt.Printf("二%b\n",a3)
	//八进制
	fmt.Printf("八%o\n",a3)
	//十六进制
	fmt.Printf("十六%x\n",a3)

	s :="我爱中国"
	b := true
	fmt.Printf("%T\n",s)
	//输出字符串
	fmt.Printf("%s\n",s)
	//输出值并且位值加上修饰符
	fmt.Printf("%#v\n",s)
	fmt.Printf("%#v\n",a3)
	fmt.Printf("%#v\n",b)
	//输出单个字符
	for i:= 0;i< len(s);i++  {
		fmt.Printf("%c",s[i])
	}
   fmt.Println(" ")
	for _, q:=range s {
		fmt.Printf("%c",q)
	}


}