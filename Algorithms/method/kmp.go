package method

import "fmt"

func KMPDemo() {
	str1 := "BBC ABCDAB ABCDABCDABDE"
	str2 := "ABCDABD"

	next := kmpNext(str2)
	//fmt.Println(next)

	res := kmpSearch(str1, str2, next)
	fmt.Println(res)
}

func kmpSearch(str1, str2 string, next []int) int {
	j := 0
	for i := 0; i < len(str1); i++ {
		for j > 0 && str1[i] != str2[j] {
			j = next[j-1]
		}
		if str1[i] == str2[j] {
			j++
		}
		if j == len(str2) {
			return i - j + 1
		}
	}
	return -1
}

//获取到一个字符串（子串）的部分匹配值表
func kmpNext(dest string) []int {
	num := len(dest)
	next := make([]int, num)
	next[0] = 0
	j := 0
	for i := 1; i < num; i++ {
		for j > 0 && dest[i] != dest[j] {
			j = next[j-1]
		}
		if dest[i] == dest[j] {
			j++
		}
		next[i] = j
	}
	return next
}
