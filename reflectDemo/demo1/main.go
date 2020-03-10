package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType)

	rVal := reflect.ValueOf(b)

	//获取kind
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Println(kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iv=%v iv type=%T \n", iV, iV)

	stu, ok := iV.(Student) //断言
	if ok {
		fmt.Println(stu.Name, stu.Age)
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//var num int = 100
	//reflectTest01(num)

	stu := Student{
		Name: "Tom",
		Age:  22,
	}

	reflectTest02(stu)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}
