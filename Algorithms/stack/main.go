package main

import "fmt"

type user struct {
	arr []interface{}
}

func main() {
	//arrayStack.TestStack()
	//arrayStack.Calculate()
	//intarr := []int{1,2,3}
	stringarr := []string{"aa", "BB", "dd"}
	var user1 []interface{} = make([]interface{}, 3)
	for i, d := range stringarr {
		user1[i] = d
	}
	fmt.Println(user1)
}
