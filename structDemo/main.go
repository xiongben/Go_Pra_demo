package main

import (
	"encoding/json"
	"fmt"
)

//如果结构体的字段类型为：指针，slice，map的零值为nil，即还没有分配空间，需要先make，再使用
type Cat struct {
	Name  string
	Age   int
	Color string
	ptr   *int
	slice []int
	map1  map[string]string
}

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightUp Point
}

type Rect2 struct {
	leftUp, rightUP *Point
}

type Monster struct {
	Name  string `json:"name"` //tag 结构体的标签
	Age   int    `json:"age"`
	Skill string `json:"skill"`
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

	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(r1)

	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Println(&r2.leftUp, &r2.rightUP)
	fmt.Println(r2.leftUp, r2.rightUP)
	fmt.Println(*r2.leftUp, *r2.rightUP)

	monster1 := Monster{"牛魔王", 500, "芭蕉扇"}
	//序列化为json字符串
	str1, err := json.Marshal(monster1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(str1))
}
