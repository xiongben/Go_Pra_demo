package main

import (
	"fmt"
)

//在main方法执行之前执行
func init() {
	fmt.Println("this is init function")
}

func main() {
	//var i int = 5
	//fmt.Printf("%b", i)
	//
	////8进制，0开头
	//var j int = 011
	//fmt.Println("j=", j)
	//
	////16进制，0x或0X开头表示
	//var k int = 0x11
	//fmt.Println("k=", k)
	//
	////位运算测试
	//fmt.Println(2&3, 2|3, 2^3)
	//fmt.Println(-2 & 3)
	//fmt.Println(1>>2, 1<<2)
	//
	//count.Count()
	//res1 := count.Sum(10, 1, 2, 3, 4, 5)
	//fmt.Println(res1)
	//fmt.Println(count.Age)

	//闭包问题
	//f := addUpper()
	//fmt.Println(f(1))
	//fmt.Println(f(2))

	//defer
	//res3 := count.DeferTest(2,3)
	//fmt.Println(res3)

	//num2 := new(int)
	//*num2 = 100
	//fmt.Printf("%T,%v,%v,%v",num2,num2,&num2,*num2)

	//错误处理
	test()
	print("==end=====")
}

func addUpper() func(int) int {
	var n int = 10
	return func(i int) int {
		n = n + i
		return n
	}
}

func test() int {
	defer func() {
		//捕获异常
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	return res
}
