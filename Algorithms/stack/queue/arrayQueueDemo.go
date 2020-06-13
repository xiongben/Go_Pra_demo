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
	var key string
	for {
		fmt.Println("s,show:显示队列")
		fmt.Println("e,exit:退出程序")
		fmt.Println("a,add:添加数据到队列")
		fmt.Println("g,get:从队列中取出数据")
		fmt.Println("h,head:查看队列头的数据")
		fmt.Scanf("%v\n", &key)
		switch key {
		case "s":
			arrayQueue1.showQueue()
			break
		case "a":
			fmt.Println("请输入一个数")
			var val int
			fmt.Scanf("%v\n", &val)
			arrayQueue1.addQueue(val)
			break
		case "g":
			res, err := arrayQueue1.getQueue()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("取出的数据是：", res)
			break
		case "h":
			res, err := arrayQueue1.headQueue()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("队列头部的数据是：", res)
			break
		case "e":
			goto Loop
			break
		default:
			break
		}
	}
Loop:
	fmt.Println("程序退出~~~~")
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

func (this *arrayQueue) getQueue() (val int, err error) {
	if this.isEmpty() {
		err = errors.New("队列为空，没有数据")
	}
	this.front++
	val = this.arr[this.front]
	return val, err
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

func (this *arrayQueue) headQueue() (res int, err error) {
	if this.isEmpty() {
		err = errors.New("队列为空，没有数据")
	}
	res = this.arr[this.front+1]
	return res, err
}
