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

func main() {
	var num int = 100
	reflectTest01(num)
}
