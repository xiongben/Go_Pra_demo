package tree

import "fmt"

//堆排序
func HeapSortDemo() {
	arr1 := []int{4, 6, 8, 5, 9}
	//adjustHeap(arr1,1,len(arr1))
	heapSort(arr1)
}

func heapSort(arr []int) {
	temp := 0
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, len(arr))
	}
	for j := len(arr) - 1; j > 0; j-- {
		temp = arr[j]
		arr[j] = arr[0]
		arr[0] = temp
		adjustHeap(arr, 0, j)
	}
	fmt.Println(arr)
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
