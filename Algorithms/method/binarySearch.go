package method

import "fmt"

func BinarySearchDemo() {
	arr1 := []int{1, 3, 8, 10, 11, 67, 100}
	index := binarySearch(arr1, 8)
	fmt.Println(index)
}

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
