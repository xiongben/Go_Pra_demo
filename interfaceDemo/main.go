package main

import "fmt"

type AInterface interface {
	say()
}

type Stu struct {
	Name string
}
type T interface {
}

func (s Stu) say() {
	fmt.Println("say something!")
}

func (s Stu) hello() {
	fmt.Println("this is hello!")
}
func main() {
	var stu Stu
	var a AInterface = stu
	a.say()

	//空接口可以赋值多种类型
	var aa T = 66
	aa = "ssss"
	aa = true
	fmt.Println(aa)
}
