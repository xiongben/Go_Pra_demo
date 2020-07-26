package sort

func radixSort(arr []int) {
	max := arr[0]
	arrlength := len(arr)
	//得到数组中最大的数
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	//得到最大的数是几位数
	maxLength := len(string(max))
	bucket := make([][10]int, arrlength)
	var bucketElementCounts [10]int

}
