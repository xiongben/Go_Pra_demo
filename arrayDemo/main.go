package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	//初始化
	arr2 := [...][3]int{{1, 2, 3}, {1, 2, 3}, {4, 5, 6}, {4, 5, 6}}
	fmt.Println(arr2)

	//遍历
	for i, v := range arr2 {
		fmt.Printf("i=%v v=%v", i, v)
	}
}
