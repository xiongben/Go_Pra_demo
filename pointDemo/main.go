package main

import (
	"fmt"
)

func main() {
	var i int = 100
	//ptr是一个指针变量
	var ptr *int = &i
	fmt.Println(&i)
	fmt.Println("ptr is :", ptr)
	fmt.Println("ptr的地址：", &ptr)
	fmt.Println("ptr指向的值：", *ptr)

	*ptr = 55
	fmt.Println(i)

	//从控制台接收用户信息,方式一
	var name string
	var age byte
	var sal float32
	var ispass bool
	//fmt.Println("请输入姓名")
	//fmt.Scanln(&name)
	//fmt.Println("请输入年龄")
	//fmt.Scanln(&age)
	//fmt.Println("请输入薪水")
	//fmt.Scanln(&sal)
	//fmt.Println("请输入是否通过")
	//fmt.Scanln(&ispass)
	//fmt.Printf("name is %v,age is %v,salary is %v,ispass is %v",name,age,sal,ispass)

	//方式二
	fmt.Println("请输入你的名字，年龄，薪水，是否通过，用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &ispass)
	fmt.Printf("name is %v,age is %v,salary is %v,ispass is %v", name, age, sal, ispass)
}
