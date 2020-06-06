package sort

import (
	"fmt"
	"math/rand"
	"time"
)

func TestBubbleSort() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var arr1 []int = []int{}
	for i := 0; i < 10; i++ {
		arr1 = append(arr1, r.Intn(100))
		//arr1[i] = r.Intn(100)
	}
	fmt.Println(arr1)
	Bubblesort(arr1)
	fmt.Println(arr1)
}

func Bubblesort(arr []int) {
	temp := 0
	flag := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				flag = true
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
		if !flag {
			break
		} else {
			flag = false
		}
	}
}
