package tree

import (
	"fmt"
	"strconv"
)

func BinaryTreeDemo() {
	hero1 := HeroNode{
		No:    1,
		Name:  "宋江",
		Left:  nil,
		Right: nil,
	}
	hero2 := HeroNode{
		No:    2,
		Name:  "吴用",
		Left:  nil,
		Right: nil,
	}
	hero3 := HeroNode{
		No:    3,
		Name:  "卢俊义",
		Left:  nil,
		Right: nil,
	}
	hero4 := HeroNode{
		No:    4,
		Name:  "林冲",
		Left:  nil,
		Right: nil,
	}
	hero5 := HeroNode{
		No:    5,
		Name:  "关胜",
		Left:  nil,
		Right: nil,
	}
	hero1.Left = &hero2
	hero1.Right = &hero3
	hero3.Left = &hero5
	hero3.Right = &hero4
	binaryTree1 := BinaryTree{
		root: &hero1,
	}
	//binaryTree1.preOrder()
	//fmt.Println("======")
	//binaryTree1.infixOrder()
	//fmt.Println("======")
	//binaryTree1.postOrder()

	resNode := binaryTree1.preOrderSearch(5)
	if resNode != nil {
		fmt.Println(resNode)
	}
}

type BinaryTree struct {
	root *HeroNode
}

func (this *BinaryTree) preOrder() {
	if this.root != nil {
		this.root.preOrder()
	} else {
		fmt.Println("tree is null,can't traverse!")
	}
}

func (this *BinaryTree) infixOrder() {
	if this.root != nil {
		this.root.infixOrder()
	} else {
		fmt.Println("tree is null,can't traverse!")
	}
}

func (this *BinaryTree) postOrder() {
	if this.root != nil {
		this.root.postOrder()
	} else {
		fmt.Println("tree is null,can't traverse!")
	}
}

func (this *BinaryTree) preOrderSearch(no int) *HeroNode {
	if this.root != nil {
		return this.root.preOrderSearch(no)
	} else {
		return nil
	}
}

type HeroNode struct {
	No    int
	Name  string
	Left  *HeroNode
	Right *HeroNode
}

func (this *HeroNode) String() string {
	return "HeroNode [no=" + strconv.Itoa(this.No) + ",name=" + this.Name + "]"
}

func (this *HeroNode) delNode(no int) {
	if this.Left != nil && this.Left.No == no {
		this.Left = nil
		return
	}
	if this.Right != nil && this.Right.No == no {
		this.Right = nil
		return
	}
	if this.Left != nil {
		this.Left.delNode(no)
	}
	if this.Right != nil {
		this.Right.delNode(no)
	}
}

func (this *HeroNode) preOrder() {
	fmt.Println(this)
	if this.Left != nil {
		this.Left.preOrder()
	}
	if this.Right != nil {
		this.Right.preOrder()
	}
}

func (this *HeroNode) infixOrder() {
	if this.Left != nil {
		this.Left.infixOrder()
	}
	fmt.Println(this)
	if this.Right != nil {
		this.Right.infixOrder()
	}
}

func (this *HeroNode) postOrder() {
	if this.Left != nil {
		this.Left.postOrder()
	}
	if this.Right != nil {
		this.Right.postOrder()
	}
	fmt.Println(this)
}

func (this *HeroNode) preOrderSearch(no int) *HeroNode {
	var resNode *HeroNode = nil
	if this.No == no {
		return this
	}
	if this.Left != nil {
		resNode = this.Left.preOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}
	if this.Right != nil {
		resNode = this.Right.preOrderSearch(no)
	}
	return resNode
}

func (this *HeroNode) infixOrderSearch(no int) *HeroNode {
	var resNode *HeroNode = nil
	if this.Left != nil {
		resNode = this.Left.infixOrderSearch(no)
	}

	if resNode != nil {
		return resNode
	}
	if this.No == no {
		return this
	}
	if this.Right != nil {
		resNode = this.Right.infixOrderSearch(no)
	}
	return resNode
}

func (this *HeroNode) postOrderSearch(no int) *HeroNode {
	var resNode *HeroNode = nil
	if this.Left != nil {
		resNode = this.Left.postOrderSearch(no)
	}
	if resNode != nil {
		return resNode
	}
	if this.Right != nil {
		resNode = this.Right.postOrderSearch(no)
	}
	if this.No == no {
		return this
	}

	return resNode
}
