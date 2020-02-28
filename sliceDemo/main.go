package main

import "fmt"

func main() {
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[2] = 20
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))

	var slice2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2))

	for i, v := range slice2 {
		fmt.Println(i, "====>", v)
	}

	//切片可以再切,指向同一个内存区域
	slice3 := slice2[1:3]
	fmt.Println(slice3)
	slice3[1] = 888
	fmt.Println(slice2)

	//append 动态增长
	slice4 := []int{1, 2, 3, 4, 5, 6, 7}
	slice4 = append(slice4, 400, 500, 600)
	slice4 = append(slice4, slice4...)
	fmt.Println(slice4)

	//string的底层是一个byte数组，因此可以进行切片处理操作
	//切片操作的时候，string不可变，不可修改
	str := "helloworld"
	slice5 := str[3:]
	fmt.Println(slice5)

	//修改string
	arr1 := []rune(str)
	arr1[0] = '呗'
	str2 := string(arr1)
	fmt.Println(str2)

}
