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
	//test()
	//print("==end=====")

	//数组
	//var hens [2]float64
	//hens[0] = 3.0
	//hens[1] = 5.0
	//
	//totalweight := 0.0
	//for i:=0;i<2;i++ {
	//	totalweight+= hens[i]
	//}
	//print(totalweight)

	var numArr [3]int = [3]int{1, 2, 3}
	//var numArr2 = [...]int{1,2,3,4,5}
	//fmt.Println(numArr)
	//fmt.Println(numArr2)
	//for index,value := range numArr2 {
	//	fmt.Println(index,"=>",value)
	//}

	testArr(numArr)
	fmt.Println(numArr)
	testArr2(&numArr)
	fmt.Println(numArr)

	//切片

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

func testArr(arr1 [3]int) {
	arr1[0] = 100
}

func testArr2(arr1 *[3]int) {
	(*arr1)[0] = 100
}
