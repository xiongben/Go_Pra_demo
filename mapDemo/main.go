package main

import (
	"fmt"
	"sort"
)

//1.map是引用类型
//2.map能动态增长键值对
//3.map的value经常使用struct类型

func main() {
	var a map[string]string
	a = make(map[string]string, 10)
	a["no1"] = "xiaoming"
	a["no2"] = "xiaohong"
	fmt.Println(a)

	cities := make(map[string]string)
	cities["no1"] = "beijing"
	cities["no2"] = "wuhang"
	cities["no3"] = "changsha"
	fmt.Println(cities)

	//查找
	val, ok := cities["no4"]
	fmt.Println(val, ok)

	//删除
	//delete(cities,"no1")
	//fmt.Println(cities)

	//一次性删除所有key
	//cities = make(map[string]string)
	//fmt.Println(cities)

	heroes := map[string]string{
		"hero1": "xiaoming",
		"hero2": "xiaohong",
	}
	fmt.Println(heroes)

	studentmap := make(map[string]map[string]string)
	studentmap["stu01"] = make(map[string]string, 3)
	studentmap["stu01"]["name"] = "tom"
	studentmap["stu01"]["sex"] = "man"
	studentmap["stu01"]["address"] = "beijing"

	studentmap["stu02"] = make(map[string]string, 3)
	studentmap["stu02"]["name"] = "jack"
	studentmap["stu02"]["sex"] = "man"
	studentmap["stu02"]["address"] = "beijing"

	fmt.Println(studentmap)

	//map的遍历
	//for k,v :=range cities {
	//	fmt.Println("遍历：",k,"=>",v)
	//}
	//
	//for k,v :=range studentmap {
	//	for k2,v2 := range v{
	//		fmt.Println(k,"=>",k2,"=>",v2)
	//	}
	//}

	//map切片
	//首先声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"

	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "玉兔精"
		monsters[1]["age"] = "500"

	}

	newMonster := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}

	monsters = append(monsters, newMonster)

	fmt.Println(monsters)

	//map排序
	map2 := make(map[int]int, 10)
	map2[10] = 100
	map2[1] = 13
	map2[4] = 76
	map2[3] = 66
	map2[5] = 45

	fmt.Println(map2)
	//将map的key放入切片中，对切片进行排序，遍历切片，输出map的值
	var keys []int
	for k, _ := range map2 {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _, k := range keys {
		fmt.Println(k, "=>", map2[k])
	}
}
