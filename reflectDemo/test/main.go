package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"name"`
	Age  int
}

func (stu Student) Testfn1() {
	stu.Name = "tom"
	fmt.Println("this is test function 1", stu.Name)
}

func (stu Student) Testfn2(aa string, bb int) {
	fmt.Println("this is test function 2", aa)
	fmt.Println(bb)
}

func main() {
	stu := Student{
		Name: "xiaoming",
		Age:  12,
	}

	//stu.testfn1()
	//stu.testfn2()
	//
	//fmt.Println(stu)
	rStu := reflect.ValueOf(stu)
	rType := reflect.TypeOf(stu)

	fmt.Println(rStu.NumMethod(), rStu.NumField(), rStu.Field(0))
	rStu.FieldByNameFunc(func(s string) bool {

		fmt.Println(s)
		if s == "Name" {
			fmt.Println("=====>>>>>>>======")
			return true
		}
		return false
	})

	nummethod := rType.NumMethod()
	rtypefiled1 := rType.Field(0).Tag.Get("json")
	fmt.Println(rtypefiled1, nummethod)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("rrrrrr"))
	params = append(params, reflect.ValueOf(66))
	rStu.Method(1).Call(params)
	rStu.Method(0).Call(nil)

	fnname := rType.Method(0)
	fmt.Println(fnname)

	ss := []interface{}{"ssss", 222, "tttt", true, false, 44.67, "dd"}
	fmt.Println(ss)
}
