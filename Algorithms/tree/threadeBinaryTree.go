package tree

import "fmt"

//线索化二叉树

func ThreadeBinaryTreeDemo() {
	node1 := HeroNode{No: 1, Name: "tom"}
	node2 := HeroNode{No: 3, Name: "ben"}
	node3 := HeroNode{No: 6, Name: "roll"}
	node4 := HeroNode{No: 8, Name: "amiss"}
	node5 := HeroNode{No: 10, Name: "jack"}
	node6 := HeroNode{No: 14, Name: "dim"}

	node1.Left = &node2
	node1.Right = &node3
	node2.Left = &node4
	node2.Right = &node5
	node3.Left = &node6

	threadtree1 := ThreadeBinaryTree{
		root: &node1,
		pre:  nil,
	}
	threadtree1.threadedNodes(&node1)

	//test node5
	//leftNode := node5.Left
	//rightNode := node5.Right
	//fmt.Println(leftNode,rightNode) //3,1

	threadtree1.threadeList()
}

type ThreadeBinaryTree struct {
	root *HeroNode
	pre  *HeroNode //初始化为空
}

//node是当前需要线索化的节点
func (this *ThreadeBinaryTree) threadedNodes(node *HeroNode) {
	if node == nil {
		return
	}
	this.threadedNodes(node.Left)
	if node.Left == nil {
		node.Left = this.pre
		node.leftType = 1
	}
	if this.pre != nil && this.pre.Right == nil {
		this.pre.Right = node
		this.pre.rightType = 1
	}
	this.pre = node
	this.threadedNodes(node.Right)

}

func (this *ThreadeBinaryTree) threadeList() {
	node := this.root
	for node != nil {
		for node.leftType == 0 {
			node = node.Left
		}
		fmt.Println(node)
		for node.rightType == 1 {
			node = node.Right
			fmt.Println(node)
		}
		node = node.Right
	}

}
