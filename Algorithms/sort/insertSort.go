package sort

import "fmt"

func InsertSortDemo() {
	arr1 := []int{101, 56, 34, 22, 13, 456, 33, 76, 24, 89, 4, 7}
	insertSort(arr1)
}

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		arr[insertIndex+1] = insertVal
	}
	fmt.Println(arr)
}
