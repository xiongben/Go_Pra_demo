package tree

import (
	"fmt"
	"strconv"
)

func AvlDemo() {

}

type Node2 struct {
	value int
	left  *Node2
	right *Node2
}

func (this *Node2) add(node *Node2) {
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

func (this *Node2) infixOrder() {
	if this.left != nil {
		this.left.infixOrder()
	}
	fmt.Println(this)
	if this.right != nil {
		this.right.infixOrder()
	}
}

func (this *Node2) String() string {
	return "Node [ no=" + strconv.Itoa(this.value) + " ]"
}

//返回以该节点为根节点的树的高度
func (this *Node2) height() int {
	var leftheight int
	var rightheight int
	if this.left == nil {
		leftheight = 0
	} else {
		leftheight = this.left.height()
	}
	if this.right == nil {
		rightheight = 0
	} else {
		rightheight = this.right.height()
	}
	return getmax(leftheight, rightheight)

}

func (this *Node2) search(val int) *Node2 {
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

func (this *Node2) searchParent(val int) *Node2 {
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

func getmax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
