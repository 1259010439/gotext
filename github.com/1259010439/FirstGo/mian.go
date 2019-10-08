//程序所属的包
package main
//导入语句
import "fmt"

//全局变量声明后可以不使用而局部变量声明后必须使用，和java类似Java中的局部变量也是必须赋值才行
var s1 string
var s2 int

const name  = "a"
//程序的入口
func main()  {
	//iota每增加一行新的常量声明就会加1
	const (
		d1 = iota
		d2
		_
		d3

	)

	fmt.Println(d1,d2,d3)
}

//匿名变量
