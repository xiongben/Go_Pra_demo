package linkedlist

type CircleSingleLinkedList struct {
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
