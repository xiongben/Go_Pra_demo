package sort

func quickSort(arr []int, left int, right int) {
	l := left
	r := right
	pivot := arr[l]
	temp := 0
	for l != r {
		for arr[l] < pivot && l < r {
			l += 1
		}
		for arr[r] > pivot && l < r {
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
