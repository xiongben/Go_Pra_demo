package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Color string
}

func main() {
	var cat1 Cat
	cat1.Name = "xiaobai"
	cat1.Age = 2
	cat1.Color = "white"
	fmt.Println(cat1)
}
