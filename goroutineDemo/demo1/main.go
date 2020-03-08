package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 1; i < 5; i++ {
		fmt.Println("hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test2() {
	for i := 1; i < 5; i++ {
		fmt.Println("hello,goland " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test()
	test2()
	cpuNum := runtime.NumCPU()
	fmt.Println(cpuNum)

}
