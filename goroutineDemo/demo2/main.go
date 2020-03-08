package main

import "sync"

//计算1-200各个数的阶乘，并放入到一个map中

var mymap = make(map[int]int, 10)
var lock sync.Mutex

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	lock.Lock()
	mymap[n] = res
	lock.Unlock()
}

func main() {

}
