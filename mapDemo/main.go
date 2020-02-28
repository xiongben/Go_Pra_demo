package main

import "fmt"

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

	fmt.Println(studentmap)
}
