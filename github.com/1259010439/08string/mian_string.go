package main

import (
	"fmt"
	"strings"
)

func main()  {
	//字符串的转义
	path :="E:\\java\\spring-boot\\spring-boot-freemarker\\src"
	fmt.Printf("%s\n",path)
	//单个的字符是int的类型
	x :='l'
	fmt.Printf("%T %v",x,x)
	//多行字符串反引号中的这个字符串都是原样输出的
	s2:=`
		窗前明月光
		疑是地上霜
     `
	fmt.Printf(s2)


	//字符串相关操作
	s3:="李白"
	s4 := "大诗人"
	s5 :="what is this"
	s6 :="chinese"
	//字符串的拼接 和Java一样可以用+进行拼接
	fmt.Println(s3+s4)
	//还可以用fmt包中的Sprintf这个方法进拼接这个方法可以返回一个string类型的值
	s34:= fmt.Sprintf("%s%s",s3,s4)
	fmt.Printf("%s\n",s34)
	//字符串的长度go中计算字符串的长度计算的是编码的长度，因为go语言统一用的是utf-8所有一个汉字占用的是8个字节
	n := len(s34)
	fmt.Printf("'%v %v\n",s5,len(s5))
	fmt.Printf("'%v %v\n",s6,len(s6))
	fmt.Println(n)
	//字符串的分割
	sum := strings.Split(path,"\\")
    fmt.Printf("%v %T\n",sum,sum)
    //判断是否有包含
    fmt.Println(strings.Contains(path,"java"))
	//前缀和后缀
	//前缀
	fmt.Println(strings.HasPrefix(path,"E"))
	//后缀
	fmt.Println(strings.HasSuffix(path,"E"))
	//判断位置字串所在的位置返回的是第一个出现该字符 的位置
	fmt.Println(strings.Index(path,"s"))
	//返回最后一次出现的位置
	fmt.Println(strings.LastIndex(path,"s"))
	//链接可以将string数组进行一个链接
	fmt.Println(strings.Join(sum,"~~"))

	//字符串的遍历当全是英文的时候可以用这种方法但是当字符串中出现中文或者其他的语言的时候就会乱码
	//因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字
	// 符串，否则就会出现乱码的结果。字符串底层是一个byte数组，所以可以和[]byte类型
	// 相互转换。字符串是不能修改的 字符串是由byte字节组成，所以字符串的长度是byte字节的长度。
	// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成。
	for i:=0;i<len(s6) ;i++  {
		fmt.Printf("%c\n",s6[i])
	}

	//rune类型的字符串输出
	for _, r := range s3 { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println("  ")
    //在go语言中的string都是由byte字符组成的但是在字符串中可能由中文也可能有英文，英文就是我们常见的ASCII 码（占一个字节）叫byte
    //但是中文就要（占3-4个字节）就叫rune其实本质上都是int类型的数字

	//字符串修改，在go中的字符串是不能进行修改的，但是在go中的字符串的本质呢就是int的数组（数组中存放着对应的ASCII码）所以在
	//修改的时候要转换成数组在修改
	s7 := "hello 中国"
	s8:=[]int32(s7)  //s8:=[]int32(s7)这样写也可以 这个就是相当于将string转换成了一个rune的数组rune数组其实就是int32数组
	s8[0] = '这'
	fmt.Println(s7)
	for i:=0;i< len(s8);i++  {
		fmt.Printf("%c",s8[i])
	}

}


