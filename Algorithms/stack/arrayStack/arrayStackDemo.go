package arrayStack

import (
	"errors"
	"fmt"
)

type ArrayStack struct {
	maxSize int
	stack   []int
	top     int
}

func (this *ArrayStack) isFull() bool {
	return this.top == this.maxSize-1
}

func (this *ArrayStack) isEmpty() bool {
	return this.top == -1
}

func (this *ArrayStack) push(val int) {
	if this.isFull() {
		fmt.Println("栈满")
		return
	}
	this.top++
	this.stack[this.top] = val
}

func (this *ArrayStack) pop() (val int, err error) {
	if this.isEmpty() {
		fmt.Println("栈空，没有数据")
		err = errors.New("stack is null")
		return
	}
	val = this.stack[this.top]
	this.top--
	return val, nil
}

func (this *ArrayStack) list() (err error) {
	if this.isEmpty() {
		fmt.Println("栈空，没有数据")
		err = errors.New("stack is null")
		return
	}
	for key, val := range this.stack {
		fmt.Println(key, val)
	}
	return
}
