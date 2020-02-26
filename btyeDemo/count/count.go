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
