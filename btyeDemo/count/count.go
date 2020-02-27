package count

import "fmt"

var Age int

func init() {
	Age = 15
}

func Count() {
	fmt.Println("this is function count")
}

func Sum(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func DeferTest(n1, n2 int) int {
	//将值拷贝入栈
	defer fmt.Println("this is defer test n1:", n1)
	defer fmt.Println("this is defer test n2:", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res:", res)
	return res
}
