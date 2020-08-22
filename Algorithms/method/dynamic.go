package method

import "fmt"

func DynamicPackageDemo() {
	w := []int{1, 4, 3}            //物品重量
	val := []int{1500, 3000, 2000} //物品价值
	m := 4                         //背包容量
	n := len(val)                  //物品个数，每种只能放一个

	v := make([][]int, n+1)
	for i := range v {
		v[i] = make([]int, m+1)
	}

	for i := 0; i < len(v); i++ {
		v[i][0] = 0
	}
	for j := 0; j < len(v[0]); j++ {
		v[0][j] = 0
	}

	for i := 1; i < len(v); i++ {
		for j := 1; j < len(v[0]); j++ {
			if w[i-1] > j {
				v[i][j] = v[i-1][j]
			} else {
				v[i][j] = getmax(v[i-1][j], val[i-1]+v[i-1][j-w[i-1]])
			}
		}
	}

	for i := range v {
		for j := range v[i] {
			fmt.Print(v[i][j], " ")
		}
		fmt.Println()
	}
}

func getmax(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
