package sort

import "fmt"

func QuickSortDemo() {
	arr1 := []int{4, 7, 2, 5, 77, 33, 89, 12, 3, 1}
	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
	//testfor()
}

func quickSort(arr []int, left int, right int) {
	l := left
	r := right
	if left > right {
		return
	}
	pivot := arr[left]
	temp := 0
	fmt.Println(pivot)
	for l != r {
		for arr[r] >= pivot && l < r {
			r--
		}
		for arr[l] <= pivot && l < r {
			l++
		}
		//fmt.Println("l=",l)

		fmt.Println(l, r)
		if l < r {
			temp = arr[l]
			arr[l] = arr[r]
			arr[r] = temp
		}
	}
	arr[left] = arr[l]
	arr[l] = pivot
	//fmt.Println(arr)
	quickSort(arr, left, l-1)
	quickSort(arr, l+1, right)
	return
}

func testfor() {
	i := 1
	j := 15
	k := 10
	for i < k {
		i++

	}
	for j > k {
		j--
	}
	fmt.Println(i, j)
}
