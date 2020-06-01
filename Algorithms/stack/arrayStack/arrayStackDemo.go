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

func TestStack() {
	stack := &ArrayStack{
		maxSize: 4,
		stack:   make([]int, 4),
		top:     -1,
	}
	var key string
	//loop := true

	for {
		fmt.Println("show: 表示显示栈")
		fmt.Println("exit: 退出程序")
		fmt.Println("push: 入栈")
		fmt.Println("pop: 出栈")
		fmt.Println("输入你的选择")
		fmt.Scanf("%v\n", &key)
		switch key {
		case "show":
			stack.list()
			break
		case "push":
			fmt.Println("请输入一个数")
			var value int
			fmt.Scanf("%d\n", &value)
			stack.push(value)
			break
		case "pop":
			res, err := stack.pop()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("出栈的数据是：", res)
			break
		case "exit":
			goto Loop
			break
		default:
			break
		}
	}
Loop:
	fmt.Println("程序退出~~~~")
}
