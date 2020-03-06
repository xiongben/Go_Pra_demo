package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	a = point
	//var b Point
	//b = a 会报错
	b, ok := a.(Point)
	if ok {
		fmt.Println("assert success", b)
	} else {
		fmt.Println("convert fail")
	}
	fmt.Println(b)
}
