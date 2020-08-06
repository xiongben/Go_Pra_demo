package sort

import "fmt"

func ShellSortDemo() {
	arr1 := []int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	shellSort(arr1)
}

func shellSort(arr []int) {
	//gap为步长
	for gap := len(arr) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(arr); i++ {
			for j := i - gap; j >= 0; j -= gap {
				if arr[j] > arr[j+gap] {
					temp := arr[j]
					arr[j] = arr[j+gap]
					arr[j+gap] = temp
				}
			}
		}
	}
	fmt.Println(arr)
}
