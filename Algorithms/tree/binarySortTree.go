package tree

import (
	"fmt"
	"strconv"
)

func BinarySortTreeDemo() {
	arr1 := []int{7, 3, 10, 12, 5, 1, 9}
	binarySortTree1 := BinarySortTree{root: nil}
	for _, v := range arr1 {
		node := Node{value: v}
		binarySortTree1.add(&node)
	}
	binarySortTree1.infixOrder()
}

type BinarySortTree struct {
	root *Node
}

func (this *BinarySortTree) add(node *Node) {
	if this.root == nil {
		this.root = node
	} else {
		this.root.add(node)
	}
}

func (this *BinarySortTree) infixOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("The tree is null,can't traverse")
	}
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (this *Node) add(node *Node) {
	if node == nil {
		return
	}
	if node.value < this.value {
		if this.left == nil {
			this.left = node
		} else {
			this.left.add(node)
		}
	} else {
		if this.right == nil {
			this.right = node
		} else {
			this.right.add(node)
		}
	}
}

func (this *Node) infixOrder() {
	if this.left != nil {
		this.left.infixOrder()
	}
	fmt.Println(this)
	if this.right != nil {
		this.right.infixOrder()
	}
}

func (this *Node) String() string {
	return "Node [ no=" + strconv.Itoa(this.value) + " ]"
}
