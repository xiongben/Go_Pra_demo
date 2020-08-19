package tree

import (
	"fmt"
	"strconv"
)

func AvlDemo() {
	//arr1 := []int{4, 3, 6, 5, 7, 8}
	//arr1 := []int{10,12,8,9,7,6}
	arr1 := []int{10, 11, 7, 6, 8, 9}
	avlTree1 := AVLTree{root: nil}
	for _, v := range arr1 {
		node := Node2{value: v}
		avlTree1.add(&node)
	}
	avlTree1.infixOrder()
	fmt.Println(avlTree1.root.height())
	fmt.Println(avlTree1.root.leftHeight(), avlTree1.root.rightHeight())
}

type AVLTree struct {
	root *Node2
}

func (this *AVLTree) add(node *Node2) {
	if this.root == nil {
		this.root = node
	} else {
		this.root.add(node)
	}
}

func (this *AVLTree) infixOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("The tree is null,can't traverse")
	}
}

func (this *AVLTree) search(val int) *Node2 {
	if this.root == nil {
		return nil
	} else {
		return this.root.search(val)
	}
}

//返回以node为根节点的二叉排序树的最小节点的值
//删除找到的最小节点
func (this *AVLTree) delRightTreeMin(node *Node2) int {
	target := node
	for target.left != nil {
		target = target.left
	}
	this.delNode(target.value)
	return target.value
}

func (this *AVLTree) searchParent(val int) *Node2 {
	if this.root == nil {
		return nil
	} else {
		return this.root.searchParent(val)
	}
}

func (this *AVLTree) delNode(val int) {
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
			minval := this.delRightTreeMin(targetNode.right)
			targetNode.value = minval
		} else { //删除只有一颗子树的节点
			if targetNode.left != nil {
				if parent != nil {
					if parent.left.value == val {
						parent.left = targetNode.left
					} else {
						parent.right = targetNode.left
					}
				} else {
					this.root = targetNode.left
				}

			} else {
				if parent != nil {
					if parent.left.value == val {
						parent.left = targetNode.right
					} else {
						parent.right = targetNode.right
					}
				} else {
					this.root = targetNode.right
				}

			}

		}
	}
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
	if this.rightHeight()-this.leftHeight() > 1 {
		//如果它的右子树的左子树高度大于它的右子树的右子树高度
		if this.right.leftHeight() > this.right.rightHeight() && this.right != nil {
			this.right.rightRotate()
			this.leftRotate()
		} else {
			this.leftRotate()
		}
		return //重要
	}
	if this.leftHeight()-this.rightHeight() > 1 {
		//如果它的左子树的右子树高度大于它的左子树的高度
		if this.left != nil && this.left.rightHeight() > this.left.leftHeight() {
			this.left.leftRotate()
			this.rightRotate()
		} else {
			this.rightRotate()
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

//返回左子树的高度
func (this *Node2) leftHeight() int {
	if this.left == nil {
		return 0
	}
	return this.left.height()
}

//返回右子树高度
func (this *Node2) rightHeight() int {
	if this.right == nil {
		return 0
	}
	return this.right.height()
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
	return getmax(leftheight, rightheight) + 1

}

//左旋转方法
func (this *Node2) leftRotate() {
	newNode := Node2{value: this.value}
	newNode.left = this.left
	newNode.right = this.right.left
	this.value = this.right.value
	this.right = this.right.right
	this.left = &newNode
}

//右旋转方法
func (this *Node2) rightRotate() {
	newNode := Node2{value: this.value}
	newNode.right = this.right
	newNode.left = this.left.right
	this.value = this.left.value
	this.left = this.left.left
	this.right = &newNode
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
