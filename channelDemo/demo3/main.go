package main

import "fmt"

func putNum(intchan chan int) {
	for i := 0; i < 8000; i++ {
		intchan <- i
	}
	close(intchan)
}

func primeNum(intchan chan int, primechan chan int, exitchan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intchan
		if !ok {
			break
		}
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primechan <- num
		}
	}
	exitchan <- true

}

func main() {

	intchan := make(chan int, 1000)
	primechan := make(chan int, 2000)
	exitchan := make(chan bool, 8)

	go putNum(intchan)
	for i := 0; i < 8; i++ {
		go primeNum(intchan, primechan, exitchan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitchan
		}
		close(primechan)
	}()

	for {
		res, ok := <-primechan
		if !ok {
			break
		}
		fmt.Printf("素数=%d \n", res)
	}

	fmt.Println("main主线程退出")
}
