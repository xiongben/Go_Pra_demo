package main

import (
	"fmt"
)

func main() {
	var i int = 5
	fmt.Printf("%b", i)

	//8进制，0开头
	var j int = 011
	fmt.Println("j=", j)

	//16进制，0x或0X开头表示
	var k int = 0x11
	fmt.Println("k=", k)

	//位运算测试
	fmt.Println(2&3, 2|3, 2^3)
	fmt.Println(-2 & 3)
	fmt.Println(1>>2, 1<<2)
}
