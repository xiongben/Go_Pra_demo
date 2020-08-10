package tree

//堆排序
func HeapSortDemo() {

}

func heapSort(arr []int) {

}

//将一个数组，调整成一个大顶堆
//i表示非叶子节点在数组中的索引
//length表示对多少个元素继续调整
func adjustHeap(arr []int, i int, length int) {
	temp := arr[i]
	for k := i*2 + 1; k < length; k = 2*k + 1 {
		//判断左右节点哪个最大
		if k+1 < length && arr[k] < arr[k+1] {
			k++
		}
		if arr[k] > temp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = temp
}
