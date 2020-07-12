package sort

import "fmt"

func bubbleSortTestDemo() {
	arr1 := []int{3, 9, 33, 67, -1, 56, 23, 67, 22, 12}
	bubbleSortFn(arr1)
	fmt.Println(arr1)
}

func bubbleSortFn(arr []int) {
	temp := 0
	flag := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		if !flag {
			goto Loop
		} else {
			flag = false
		}
	}
Loop:
}
