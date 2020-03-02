package main

import "fmt"

//如果结构体的字段类型为：指针，slice，map的零值为nil，即还没有分配空间，需要先make，再使用
type Cat struct {
	Name  string
	Age   int
	Color string
	ptr   *int
	slice []int
	map1  map[string]string
}

func main() {
	//var aa int32 = 100
	bb := 200
	var cat1 Cat
	cat1.Name = "xiaobai"
	cat1.Age = 2
	cat1.Color = "white"
	//先make
	cat1.slice = make([]int, 10)
	cat1.slice[0] = 100

	cat1.map1 = make(map[string]string)
	cat1.map1["name"] = "xxxx"

	cat1.ptr = new(int)
	cat1.ptr = &bb

	fmt.Println(cat1)
	//方式2
	cat2 := Cat{}
	fmt.Println(cat2)
	//方式3
	var cat3 *Cat = new(Cat)
	(*cat3).Name = "xiaoxiao" //标准写法
	cat3.Age = 22             //底层会进行处理
	fmt.Println(*cat3)

	//方式4
	var cat4 *Cat = &Cat{}
	(*cat4).Name = "Tom"
	cat4.Age = 16
	fmt.Println(*cat4)
}
