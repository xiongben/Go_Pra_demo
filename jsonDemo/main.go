package main

import (
	"encoding/json"
	"fmt"
)

//将结构体，map，切片进行序列化
type Monster struct {
	Name  string `json:"name"` //反射机制
	Age   int    `json:"age"`
	Skill string
}

func testStruct() {
	monster := Monster{
		Name:  "niumowang",
		Age:   100,
		Skill: "niumoquaneeee",
	}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(string(data))
	var monster2 Monster
	err2 := json.Unmarshal([]byte(data), &monster2) //反序列化
	if err2 != nil {
		fmt.Println("err=", err2)
	}
	fmt.Println(monster2)
}

func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})
	a["name"] = "xiaoming"
	a["age"] = 14
	a["address"] = "street 11"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(string(data))
}

func testSlice() {
	var slice []map[string]interface{}
	m1 := make(map[string]interface{})
	m1["name"] = "xiaohong"
	m1["age"] = 12
	m1["addree"] = "street 11"
	slice = append(slice, m1)

	m2 := make(map[string]interface{})
	m2["name"] = "xiaohongaaaa"
	m2["age"] = 12
	m2["addree"] = "street 11"
	slice = append(slice, m2)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println(string(data))
}

func main() {

	testStruct()
	testMap()
	testSlice()
}
