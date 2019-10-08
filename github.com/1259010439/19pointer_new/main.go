package main

import "fmt"

func main()  {

	/*
	   指针  pointer go中的指针就是一个引用
	*/
	//&取地址
	//*根据地址取值
	s1 :=[]string{"1","2","3"}
	z := &s1
	z1 :=&z
	z2 :=&z1
	fmt.Printf("%T\n",z2)
	fmt.Printf("%T\n",z1)
	fmt.Printf("%T\n",z)
	s2 := *z
	fmt.Println(s2)

    /*
     new 关键字的用法
    */
    var a1 *int//这个时候没有分配给指针内存地址所以a1指向的是空
    fmt.Println(a1)//所以这个时候输出的a1就是nil
    q :=12
    a1 =&q//这里给a1附上了一个值也就是q的指针
	fmt.Printf("地址%v 地址对应的值%d\n",a1,*a1)//这个时候输出的a1就是q的地址 ,对a1取值*a1也就能取到了
	var a2 = new(int)//new 关键字呢就是开辟一块内存地址出来
	fmt.Printf("%v\n",a2) //因为这块地址已经用new开辟出来了所以这个时候已a2经有地址所以可以输出出来了 0xc0000720c0
	fmt.Printf("%v\n",*a2)//但是因为这个时候并没有给这个地址赋值所以他的值位int类型的默认值0
	*a2 = 100//给地址赋值
	fmt.Printf("%v\n",*a2)//这个时候就能根据地址取到100这个值了

	/*
	  make和new 的区别
	  1.他们都是用来申请内存空间的
	  2.new用的地方很少，一般都是给基本数据类型来申请内存，(string,int)返回值呢是基本数据类型的指针(*string，*int)
	  3.make是用来给slice(切片)，map，chan申请内存的，make关键字返回的呢都是这些类型的本身
	 */


}
