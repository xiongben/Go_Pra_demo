package linkedlist

import "fmt"

type CircleSingleLinkedList struct {
	first *Boy
}

func (this *CircleSingleLinkedList) addBoy(nums int) {
	if nums < 1 {
		fmt.Println("输入的值不正确")
		return
	}
	var curBoy *Boy
	for i := 1; i <= nums; i++ {
		boy := Boy{
			no:   i,
			next: nil,
		}
		if i == 1 {
			this.first = &boy
			this.first.setNext(this.first)
			curBoy = this.first
		} else {
			curBoy.setNext(&boy)
			boy.setNext(this.first)
			curBoy = &boy
		}
	}
}

func (this *CircleSingleLinkedList) showBoy() {
	if this.first == nil {
		fmt.Println("没有任何小孩")
		return
	}
	curBoy := this.first
	for {
		fmt.Printf("小孩的编号是%v \n", curBoy.getNo())
		if curBoy.getNext() == this.first {
			goto Loop
		}
		curBoy = curBoy.getNext()
	}
Loop:
}

func (this *CircleSingleLinkedList) countBoy(startNo int, countNum int, nums int) {

}

type Boy struct {
	no   int
	next *Boy
}

func (this *Boy) getNo() int {
	return this.no
}

func (this *Boy) setNo(no int) {
	this.no = no
}

func (this *Boy) getNext() *Boy {
	return this.next
}
func (this *Boy) setNext(next *Boy) {
	this.next = next
}
