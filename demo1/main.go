package main

import (
	"fmt"
)

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Println(n1)
	var num1 int64 = 999999
	//按溢出处理
	var num2 int8 = int8(num1)
	fmt.Println(num2)
}
