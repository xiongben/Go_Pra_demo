package main

import (
	"fmt"
	"time"
)

func writeData(intchan chan int) {
	for i := 0; i < 10; i++ {
		intchan <- i
		fmt.Println("write data, ", i)
		time.Sleep(time.Second)
	}
	close(intchan)
}

func readData(intchan chan int, exitchan chan bool) {
	for {
		v, ok := <-intchan
		if !ok {
			fmt.Println(ok)
			break
		}
		time.Sleep(time.Second)
		fmt.Printf("readData read the data = %v \n", v)
	}
	exitchan <- true
	close(exitchan)
}

func main() {
	intchan := make(chan int, 10)
	exitchan := make(chan bool, 1)

	go writeData(intchan)
	go readData(intchan, exitchan)
	for {
		_, ok := <-exitchan
		if !ok {
			break
		}
	}
}
