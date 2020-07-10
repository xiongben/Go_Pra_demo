package linkedlist

import "fmt"

func JosefuDemo() {
	circleSingleLinkedList := CircleSingleLinkedList{}
	circleSingleLinkedList.addBoy(6)
	circleSingleLinkedList.showBoy()
	circleSingleLinkedList.countBoy(1, 3, 6)
}

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
	if this.first == nil || startNo < 1 || startNo > nums {
		fmt.Println("输入的参数有误")
		return
	}
	helper := this.first
	for {
		if helper.getNext() == this.first {
			goto Loop1
		}
		helper = helper.getNext()
	}
Loop1:
	for i := 0; i < startNo-1; i++ {
		this.first = this.first.getNext()
		helper = helper.getNext()
	}
	for {
		if helper == this.first {
			goto Loop2
		}
		for j := 0; j < countNum; j++ {
			this.first = this.first.getNext()
			helper = helper.getNext()
		}
		fmt.Printf("小孩%v出圈\n", this.first.getNo())
		this.first = this.first.getNext()
		helper.setNext(this.first)
	}
Loop2:
	fmt.Printf("the last child number is %v \n", this.first.getNo())

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
