package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int32 = 100
	var n1 float32 = float32(i)
	fmt.Println(n1)
	var num1 int64 = 999999
	//按溢出处理
	var num2 int8 = int8(num1)
	fmt.Println(num2)
	//数据类型转换成字符串
	var str1 string = fmt.Sprintf("%d", i)
	//方式一
	fmt.Printf("str type %T str=%q\n", str1, str1)
	//方式二 strconv
	var str2 string = strconv.FormatInt(int64(i), 2)
	fmt.Printf("str type %T str=%q\n", str2, str2)
}
