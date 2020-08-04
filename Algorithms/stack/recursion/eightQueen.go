package recursion

import (
	"fmt"
	"math"
)

type Queue8 struct {
	max    int
	resarr []int
}

func (this *Queue8) printResArr() {
	for i := 0; i < len(this.resarr); i++ {
		fmt.Print(this.resarr[i], " ")
	}
	fmt.Println("")
}

//n表示第几个皇后
func (this *Queue8) judge(n int) bool {
	for i := 0; i < n; i++ {
		if this.resarr[i] == this.resarr[n] || math.Abs(float64(n-i)) == math.Abs(float64(this.resarr[n]-this.resarr[i])) {
			return false
		}
	}
	return true
}
