package main

import (
	"fmt"
)

func main() {
	//使用select可以解决从管道取数据的阻塞问题

	intchan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intchan <- i
	}

	stringchan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringchan <- "hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intchan:
			fmt.Printf("从intChan读取的数据%v \n", v)
		case v := <-stringchan:
			fmt.Printf("从stringchan获取数据%v \n", v)
		default:
			fmt.Printf("取不到，不玩了===")
			return
		}
	}
}
