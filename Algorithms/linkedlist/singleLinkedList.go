package linkedlist

import (
	"fmt"
	"strconv"
)

func SingleLinkedListTestDemo() {
	hero1 := HeroNode{
		no:       1,
		name:     "卡卡西",
		nickname: "拷贝忍者",
		next:     nil,
	}
	hero2 := HeroNode{
		no:       2,
		name:     "猿飞日斩",
		nickname: "火影三代目",
		next:     nil,
	}
	hero3 := HeroNode{
		no:       3,
		name:     "宇智波鼬",
		nickname: "暗部组长",
		next:     nil,
	}
	hero4 := HeroNode{
		no:       4,
		name:     "漩涡长门",
		nickname: "晓组织首领",
		next:     nil,
	}
	hero8 := HeroNode{
		no:       8,
		name:     "凯",
		nickname: "木叶旋风",
		next:     nil,
	}
	hero12 := HeroNode{
		no:       12,
		name:     "大蛇丸",
		nickname: "科研大佬",
		next:     nil,
	}
	singleLinkedList := SingleLinkedList{}
	singleLinkedList.head = &HeroNode{
		no:       0,
		name:     "",
		nickname: "",
		next:     nil,
	}
	singleLinkedList.add(&hero1)
	singleLinkedList.add(&hero2)
	singleLinkedList.add(&hero3)
	singleLinkedList.add(&hero4)
	singleLinkedList.add(&hero8)
	singleLinkedList.add(&hero12)

	singleLinkedList.list()
	fmt.Println("=============")
	reverseList(singleLinkedList.head)
	singleLinkedList.list()
	//fmt.Println("===========")
	//singleLinkedList.del(1)
	//singleLinkedList.list()

	//fmt.Println("===========")
	//singleLinkedList.addByOrder(&hero8)
	//singleLinkedList.list()
}

//反转单链表
func reverseList(head *HeroNode) {
	if head.next == nil || head.next.next == nil {
		return
	}
	cur := head.next
	var next *HeroNode
	reverseHead := HeroNode{
		no:       0,
		name:     "",
		nickname: "",
		next:     nil,
	}
	for {
		if cur == nil {
			goto Loop
		}
		next = cur.next
		cur.next = reverseHead.next
		reverseHead.next = cur
		cur = next
	}
Loop:
	head.next = reverseHead.next
}

type SingleLinkedList struct {
	head *HeroNode
}

func (this *SingleLinkedList) getHead() *HeroNode {
	return this.head
}

func (this *SingleLinkedList) add(heronode *HeroNode) {
	temp := this.head
	for {
		if temp.next == nil {
			goto Loop
		}
		temp = temp.next
	}
Loop:
	temp.next = heronode
}

func (this *SingleLinkedList) addByOrder(heroNode *HeroNode) {
	temp := this.head
	flag := false
	for {
		if temp.next == nil {
			goto Loop
		}
		if temp.next.no > heroNode.no {
			goto Loop
		} else if temp.next.no == heroNode.no {
			flag = true
			goto Loop
		}
		temp = temp.next
	}
Loop:
	if flag {
		fmt.Println("准备插入的编号已经存在，不能加入")
	} else {
		heroNode.next = temp.next
		temp.next = heroNode
	}
}

func (this *SingleLinkedList) update(newHeroNode *HeroNode) {
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

func (this *SingleLinkedList) del(no int) {
	temp := this.head
	flag := false
	for {
		if temp == nil {
			goto Loop
		}
		if temp.next.no == no {
			flag = true
			goto Loop
		}
		temp = temp.next
	}
Loop:
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Printf("没有找到编号%v的结点，不能删除\n", no)
	}
}

func (this *SingleLinkedList) list() {
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

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode
}

func (this *HeroNode) toString() string {
	return "HeroNode [no=" + strconv.Itoa(this.no) + ", name=" + this.name + ", nickname=" + this.nickname + "]"
}
