package queue

import (
	"errors"
	"fmt"
)

func ArrayQueueDemo() {

}

type arrayQueue struct {
	maxSize int
	front   int
	rear    int
	arr     []int
}

func (this *arrayQueue) isFull() bool {
	return this.rear == this.maxSize-1
}

func (this *arrayQueue) isEmpty() bool {
	return this.rear == this.front
}

func (this *arrayQueue) addQueue(n int) {
	if this.isFull() {
		fmt.Println("队列满，不能加入数据")
		return
	}
	this.rear++
	this.arr[this.rear] = n
}

func (this *arrayQueue) getQueue() int {
	if this.isEmpty() {
		errors.New("队列为空，没有数据")
	}
	this.front++
	return this.arr[this.front]
}

func (this *arrayQueue) showQueue() {
	if this.isEmpty() {
		fmt.Println("队列为空，没有数据")
		return
	}
	for _, v := range this.arr {
		fmt.Println(v)
	}
}

func (this *arrayQueue) headQueue() int {
	if this.isEmpty() {
		errors.New("队列为空，没有数据")
	}
	return this.arr[this.front+1]
}
