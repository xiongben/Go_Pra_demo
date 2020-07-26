package main

import "fmt"

func main() {
	//数组
	//var a = [...]int{1,2,3,4}
	//var b = &a
	//fmt.Println(a[0],b[0])
	//for i,v := range b {
	//	fmt.Println(i,v)
	//}

	//var time [5][2]int
	//time[1][1] = 6
	//fmt.Println(time)
	//for range time {
	//	fmt.Println("HHH")
	//}

	var str1 = "hello"
	for i, _ := range str1 {
		fmt.Println(i, str1[:i], str1[i:])
	}
}
