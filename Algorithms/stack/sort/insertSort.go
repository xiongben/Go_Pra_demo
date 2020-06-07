package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func insertSort(arr []int) {
	insertVal := 0
	insertIndex := 0
	for i := 1; i < len(arr); i++ {
		insertVal = arr[i]
		insertIndex = i - 1
		for insertIndex >= 0 && insertVal < arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}
}

func TestInsertSort() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var arr1 []int = []int{}
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, r.Intn(100))
		//arr1[i] = r.Intn(100)
	}
	fmt.Println(arr1)
	insertSort(arr1)
	fmt.Println(arr1)
}
