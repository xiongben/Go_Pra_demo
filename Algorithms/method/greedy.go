package method

import "fmt"

func GreedyDemo() {
	broadCasts := make(map[string][]string)

	hashArr1 := []string{"beijing", "shanghai", "tianjing"}
	hashArr2 := []string{"guanzhou", "beijing", "shenzhen"}
	hashArr3 := []string{"chengdu", "shanghai", "hangzhou"}
	hashArr4 := []string{"shanghai", "tianjing"}
	hashArr5 := []string{"hangzhou", "daliang"}

	broadCasts["key1"] = hashArr1
	broadCasts["key2"] = hashArr2
	broadCasts["key3"] = hashArr3
	broadCasts["key4"] = hashArr4
	broadCasts["key5"] = hashArr5

	allAreas := []string{"beijing", "shanghai", "tianjing", "hangzhou", "daliang", "chengdu", "guanzhou", "shenzhen"}
	selects := make([]string, 0) //存放选择的电台集合

	//临时集合，存放遍历过程中电台覆盖的地区和当前没有覆盖地区的交集
	tempSet := make([]string, 0)

	var maxKey string

	for len(allAreas) != 0 {
		maxKey = ""
		for k, v := range broadCasts {
			tempSet = []string{}
			tempSet = append(tempSet, v...)
			tempSet = intersect(tempSet, allAreas)
			if len(tempSet) > 0 && (len(maxKey) == 0 || len(tempSet) > len(broadCasts[maxKey])) {
				maxKey = k
			}
		}
		if len(maxKey) != 0 {
			selects = append(selects, maxKey)
			allAreas = difference(allAreas, broadCasts[maxKey])
		}
	}

	fmt.Println(selects)
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1 - 并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
