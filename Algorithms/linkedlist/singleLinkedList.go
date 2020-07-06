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

	singleLinkedList.list()

}

type SingleLinkedList struct {
	head *HeroNode
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
