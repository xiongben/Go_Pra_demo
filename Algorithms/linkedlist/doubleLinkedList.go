package linkedlist

import (
	"fmt"
	"strconv"
)

func DoubleLinkedListTestDemo() {
	hero1 := HeroNode2{
		no:       1,
		name:     "卡卡西",
		nickname: "拷贝忍者",
	}
	hero2 := HeroNode2{
		no:       2,
		name:     "猿飞日斩",
		nickname: "火影三代目",
	}
	hero3 := HeroNode2{
		no:       3,
		name:     "宇智波鼬",
		nickname: "暗部组长",
	}
	hero4 := HeroNode2{
		no:       4,
		name:     "漩涡长门",
		nickname: "晓组织首领",
	}
	hero8 := HeroNode2{
		no:       8,
		name:     "凯",
		nickname: "木叶旋风",
	}
	hero12 := HeroNode2{
		no:       12,
		name:     "大蛇丸",
		nickname: "科研大佬",
	}
	doubleLinkedList := DoubleLinkedList{}
	doubleLinkedList.head = &HeroNode2{
		no:       0,
		name:     "",
		nickname: "",
		next:     nil,
		pre:      nil,
	}
	doubleLinkedList.addNode(&hero1)
	doubleLinkedList.addNode(&hero2)
	doubleLinkedList.addNode(&hero3)
	doubleLinkedList.addNode(&hero4)
	doubleLinkedList.addNode(&hero8)
	doubleLinkedList.addNode(&hero12)

	doubleLinkedList.list()
	fmt.Println("=============")

	doubleLinkedList.delNode(2)
	doubleLinkedList.list()

	//fmt.Println("===========")
	//singleLinkedList.addByOrder(&hero8)
	//singleLinkedList.list()
}

type DoubleLinkedList struct {
	head *HeroNode2
}

func (this *DoubleLinkedList) getHead() *HeroNode2 {
	return this.head
}

func (this *DoubleLinkedList) list() {
	if this.head.next == nil {
		fmt.Println("链表为空")
		return
	}
	temp := this.head.next
	for {
		if temp == nil {
			goto Loop
		}
		fmt.Println(temp)
		temp = temp.next
	}
Loop:
}

func (this *DoubleLinkedList) addNode(heronode *HeroNode2) {
	temp := this.head
	for {
		if temp.next == nil {
			goto Loop
		}
		temp = temp.next
	}
Loop:
	temp.next = heronode
	heronode.pre = temp
}

func (this *DoubleLinkedList) delNode(no int) {
	if this.head.next == nil {
		fmt.Println("链表为空")
		return
	}
	temp := this.head
	flag := false
	for {
		if temp == nil {
			goto Loop
		}
		if temp.no == no {
			flag = true
			goto Loop
		}
		temp = temp.next
	}
Loop:
	if flag {
		temp.pre.next = temp.next
		if temp.next != nil {
			temp.next.pre = temp.pre
		}
	} else {
		fmt.Printf("没有找到编号%v的结点，不能删除\n", no)
	}
}

func (this *DoubleLinkedList) upDate(newHeroNode *HeroNode2) {
	if this.head.next == nil {
		fmt.Println("链表为空！")
		return
	}
	temp := this.head.next
	flag := false
	for {
		if temp == nil {
			goto Loop
		}
		if temp.no == newHeroNode.no {
			flag = true
			goto Loop
		}
		temp = temp.next
	}
Loop:
	if flag {
		temp.name = newHeroNode.name
		temp.nickname = newHeroNode.nickname
	} else {
		fmt.Printf("没有找到编号%v的结果，不能修改\n", newHeroNode.no)
	}
}

type HeroNode2 struct {
	no       int
	name     string
	nickname string
	next     *HeroNode2
	pre      *HeroNode2
}

func (this *HeroNode2) toString() string {
	return "HeroNode [no=" + strconv.Itoa(this.no) + ", name=" + this.name + ", nickname=" + this.nickname + "]"
}
