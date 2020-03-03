package main

import "fmt"

type A struct {
	Num    int
	Name   string
	Radius float64
}

func (a A) test() {
	a.Name = "jack"
	fmt.Println(a.Num, a.Name)
}

func (a A) speak() {
	fmt.Println(a.Name, "is a good man")
}

func (a A) calculate(n int) {
	n = n + 100
	fmt.Println(n)
}

//为了提高效率，通常我们方法和结构体的指针类型绑定
func (a *A) area() float64 {
	return 3.14 * a.Radius * a.Radius
}

func main() {
	var a A
	a.Num = 100
	a.Name = "Tom"
	a.Radius = 5.0
	a.test()
	fmt.Println(a.Name)
	a.speak()
	a.calculate(100)

	res := (&a).area()
	fmt.Println(res)
}
