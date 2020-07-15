package sort

import "fmt"

func QuickSortDemo() {
	arr1 := []int{23, 45, 12, 66, 24, 89, 54, 11, 80, 2, 5, 77, 32}
	quickSort(arr1, 0, 12)
	fmt.Println(arr1)
}

func quickSort(arr []int, left int, right int) {
	l := left
	r := right
	pivot := arr[l]
	temp := 0
	for l != r {
		if arr[l] < pivot && l < r {
			l += 1
		}
		if arr[r] > pivot && l < r {
			r -= 1
		}
		if l < r {
			temp = arr[l]
			arr[l] = arr[r]
			arr[r] = temp
		}
	}
	arr[left] = arr[l]
	arr[l] = pivot
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
}
