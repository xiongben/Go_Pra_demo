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

func (this *BinarySortTree) search(val int) *Node {
	if this.root == nil {
		return nil
	} else {
		return this.root.search(val)
	}
}

func (this *BinarySortTree) searchParent(val int) *Node {
	if this.root == nil {
		return nil
	} else {
		return this.root.searchParent(val)
	}
}

func (this *BinarySortTree) delNode(val int) {
	if this.root == nil {
		return
	} else {
		targetNode := this.search(val)
		if targetNode == nil {
			return
		}
		if this.root.left == nil && this.root.right == nil {
			this.root = nil
			return
		}
		parent := this.searchParent(val)
		//要删除的节点是叶子节点
		if targetNode.left == nil && targetNode.right == nil {
			//判断目标节点是左节点还是右节点
			if parent.left != nil && parent.left.value == val {
				parent.left = nil
			} else if parent.right != nil && parent.right.value == val {
				parent.right = nil
			}
		} else if targetNode.left != nil && targetNode.right != nil {

		} else { //删除只有一颗子树的节点

		}
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

func (this *Node) search(val int) *Node {
	if this.value == val {
		return this
	} else if val < this.value {
		if this.left == nil {
			return nil
		}
		return this.left.search(val)
	} else {
		if this.right == nil {
			return nil
		}
		return this.right.search(val)
	}
}

func (this *Node) searchParent(val int) *Node {
	if (this.left != nil && this.left.value == val) || (this.right != nil && this.right.value == val) {
		return this
	} else {
		if val < this.value && this.left != nil {
			return this.left.searchParent(val)
		} else if val >= this.value && this.right != nil {
			return this.right.searchParent(val)
		} else {
			return nil
		}
	}
}
