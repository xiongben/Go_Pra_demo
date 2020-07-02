package linkedlist

import "strconv"

type SingleLinkedList struct {
	head HeroNode
}

func (this *SingleLinkedList) add(heronode HeroNode) {
	var temp HeroNode = this.head

}

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     HeroNode
}

func (this *HeroNode) toString() string {
	return "HeroNode [no=" + strconv.Itoa(this.no) + ", name=" + this.name + ", nickname=" + this.nickname + "]"
}
