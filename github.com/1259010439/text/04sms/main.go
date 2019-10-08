package main

import (
	"fmt"
)
var allStudent =make(map[int64]student,50)

type student struct{
	id int64
	name string
}
//查看全部学生的方法
func showStudent()  {
	//遍历map集合
	for k,v:= range allStudent{
         fmt.Printf("学号：%v  姓名：%v\n",k,v.name)
	}
}
//添加学生的方法
func addStudent() {
	//用户输入学号和姓名
	var id int64
	var name string
	fmt.Print("请输要填加学生的学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入添加学生的姓名：")
	fmt.Scanln(&name)
	//构造一个类
	var xs = student{id,name}
	allStudent[id] = xs
	fmt.Println("====添加成功===================")
}

func deleteStudent()  {
	showStudent()
	var id  int64
	fmt.Print("请输入要删除学生的学号:")
	fmt.Scanln(&id)
	delete(allStudent,id)
	fmt.Print("=================删除成功了==================")
}


func main()  {
	for i:=1;i>0 ;i=1  {
		fmt.Println("欢迎来到浩哥的学生管理系统")
		fmt.Println(`
        1.查看全部学生
        2.添加学生
        3.删除学生
        4.返回首页
   `)
		fmt.Print("请输入您要选择的操作:")
		var choose int
		fmt.Scanln(&choose)
		switch  {
		case choose ==1: showStudent()
		case choose ==2: addStudent()
		case choose ==3: deleteStudent()
		case choose ==4:continue
		default:
			fmt.Println("请正确输入")

		}
	}

}

