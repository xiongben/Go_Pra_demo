package arrayStack

import (
	"errors"
	"fmt"
)

type ArrayStack2 struct {
	maxSize int
	stack   []int
	top     int
}

func (this *ArrayStack2) peek() int {
	return this.stack[this.top]
}

func (this *ArrayStack2) isFull() bool {
	return this.top == this.maxSize-1
}

func (this *ArrayStack2) isEmpty() bool {
	return this.top == -1
}

func (this *ArrayStack2) push(val int) {
	if this.isFull() {
		fmt.Println("栈满")
		return
	}
	this.top++
	this.stack[this.top] = val
}

func (this *ArrayStack2) pop() (val int, err error) {
	if this.isEmpty() {
		fmt.Println("栈空，没有数据")
		err = errors.New("stack is null")
		return
	}
	val = this.stack[this.top]
	this.top--
	return val, nil
}

func (this *ArrayStack2) list() (err error) {
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

func (this *ArrayStack2) priority(oper int) int {
	if oper == '*' || oper == '/' {
		return 1
	} else if oper == '+' || oper == '-' {
		return 0
	} else {
		return -1
	}
}

func (this *ArrayStack2) isoper(val string) bool {
	return val == "+" || val == "-" || val == "*" || val == "/"
}

func (this *ArrayStack2) cal(num1 int, num2 int, oper string) int {
	res := 0
	switch oper {
	case "+":
		res = num1 + num2
		break
	case "-":
		res = num2 - num1
		break
	case "*":
		res = num1 * num2
		break
	case "/":
		res = num2 / num1
		break
	default:
		break
	}
	return res
}

func Calculate() {
	expression := "7*2*2-5+1-5+3-4"
	numStack := ArrayStack2{
		maxSize: 10,
		stack:   make([]int, 10),
		top:     -1,
	}
	operStack := ArrayStack2{
		maxSize: 10,
		stack:   make([]int, 10),
		top:     -1,
	}
	index := 0
	num1 := 0
	num2 := 0
	oper := 0
	res := 0
	ch := " "
	keepNum := "" //用于拼接多位数
}
