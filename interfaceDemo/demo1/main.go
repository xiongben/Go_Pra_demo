package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

//实现interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

//从小到大排列
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	//temp := hs[i]
	//hs[i] = hs[j]
	//hs[j] = temp
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	//var intSlice = []int{0, -1, 3, 5, 2, 8, 6, 60, 34}
	//fmt.Println(intSlice)

	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: "heroxxx",
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	//排序前
	for _, v := range heros {
		fmt.Println(v)
	}

	sort.Sort(heros)
	fmt.Println("=====================")
	//排序后
	for _, v := range heros {
		fmt.Println(v)
	}
}
