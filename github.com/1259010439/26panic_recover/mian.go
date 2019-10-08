package main

import "fmt"

/*
  panic和recover go中做错误处理的  他和java中的try catch不同他不是异常处理 而是异常恢复，就已经出现的错误进行挽回
  panic翻译成中文就混乱的意思 他的意思就是程序出现了错误会将错误进行一个返回，recover就会接受这个返回值然后然后可以根据
  返回的值了类型进行修复操作

  1.recover必须和panic一起使用
  2。defer语句一定要在有可能引发panic的语句之前使用
 */
func f1()  {
	fmt.Println("1")
}
func one()  {

	defer func() {
		if err:=recover();err!=nil{
                 fmt.Println("修改错误")
		}
	}()
	panic("出现了错误")

}
func f2()  {
	fmt.Println(2)
}
func main()  {
	f1()
	one()
	f2()
}
