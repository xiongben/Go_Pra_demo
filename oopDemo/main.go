package main

import "fmt"

type Student struct {
	name   string
	gender string
}

func (stu Student) getStudent() {
	fmt.Println(stu.name, stu.gender)
}

func main() {
	stu := Student{"xiaoming", "male"}
	stu.getStudent()

	stu2 := Student{
		name:   "xiaohong",
		gender: "female",
	}
	stu2.getStudent()

	var stu3 *Student = &Student{"xiaohai", "male"}
	//stu3.getStudent()
	(*stu3).getStudent()
}
