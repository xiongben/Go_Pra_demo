package sort

import "fmt"

func SelectSortDemo() {
	arr1 := []int{101, 56, 34, 22, 13, 456, 33, 76, 24, 89, 4, 7}
	selectSort(arr1)
}

func selectSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex] = arr[i]
			arr[i] = min
		}
	}
	fmt.Println(arr)
}
