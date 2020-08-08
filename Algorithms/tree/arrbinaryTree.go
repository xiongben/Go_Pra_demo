package tree

import "fmt"

//顺序存储二叉树
func ArrBinaryTreeDemo() {
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	arrbin1 := ArrBinaryTree{arr: arr1}
	arrbin1.preOrder(0)
}

type ArrBinaryTree struct {
	arr []int
}

func (this *ArrBinaryTree) preOrder(index int) {
	if this.arr == nil || len(this.arr) == 0 {
		fmt.Println("arr is null,can't traverse")
	}
	fmt.Println(this.arr[index])
	if (index*2 + 1) < len(this.arr) {
		this.preOrder(2*index + 1)
	}
	if (index*2 + 2) < len(this.arr) {
		this.preOrder(2*index + 2)
	}
}
