package queue

import (
	"errors"
	"fmt"
)

func ArrayQueueDemo() {
	arrayQueue1 := arrayQueue{
		maxSize: 3,
		front:   -1,
		rear:    -1,
		arr:     make([]int, 3),
	}
	key := ""
	for {
		fmt.Println("s,show:显示队列")
		fmt.Println("e,exit:退出程序")
		fmt.Println("a,add:添加数据到队列")
		fmt.Println("g,get:从队列中取出数据")
		fmt.Println("h,head:查看队列头的数据")
	}
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
