package recursion

import (
	"fmt"
	"math"
)

func Queue8Demo() {
	queue8t := Queue8{
		max:    8,
		resarr: nil,
		count:  0,
	}
	queue8t.resarr = make([]int, queue8t.max)
	queue8t.check(0)
	fmt.Println("total number : ", queue8t.count)
}

type Queue8 struct {
	max    int
	resarr []int
	count  int
}

func (this *Queue8) printResArr() {
	for i := 0; i < len(this.resarr); i++ {
		fmt.Print(this.resarr[i], " ")
	}
	this.count++
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

func (this *Queue8) check(n int) {
	if n == this.max {
		this.printResArr()
		return
	}
	for i := 0; i < this.max; i++ {
		//先把当前皇后n，放到该行的第一列
		this.resarr[n] = i
		if this.judge(n) {
			this.check(n + 1)
		}
	}
}
