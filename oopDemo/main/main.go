package main

import (
	"awesomeProject1/oopDemo/model"
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (s *Student) showInfo() {
	fmt.Printf("学生名-%v 年龄=%v score=%v \n", s.Name, s.Age, s.Score)
}

func (s *Student) setScore(score int) {
	s.Score = score
}

//继承
type Pupil struct {
	Student
}

type Graduate struct {
	s Student
	int
}

func (p *Pupil) testing() {
	fmt.Println("student is in testing")
}

func main() {
	fmt.Println("======oop======")
	p := model.NewPerson("jack")
	fmt.Println(p)
	p.SetAge(20)
	p.SetSal(6000.0)
	fmt.Println(p.GetAge(), p.GetSal())
	fmt.Println("======oop======")

	//方式一
	pupil := &Pupil{}
	pupil.Student.Name = "jack"
	pupil.Student.Age = 12
	pupil.Student.setScore(88)
	(*pupil).Student.showInfo()
	pupil.testing()

	//方式二
	pupil2 := &Pupil{}
	pupil2.Name = "markcle"
	pupil2.Age = 12
	pupil2.setScore(98)
	(*pupil2).showInfo()

	//方式三
	pupil3 := &Pupil{Student{"kaka", 12, 78}}
	(*pupil3).showInfo()

	fmt.Println("======oop======")
	//非匿名
	gradu := &Graduate{}
	gradu.s.Name = "kakaxi"
	gradu.s.Age = 12
	gradu.s.setScore(98)
	gradu.s.showInfo()
	gradu.int = 20
	fmt.Println(gradu)
}
