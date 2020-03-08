package main

import "fmt"

func main() {
	var intchan chan int
	intchan = make(chan int, 10)
	fmt.Printf("inchant's value is %v,inchant self address is %p \n", intchan, &intchan)

	//往管道中写入数据
	intchan <- 10
	num := 211
	intchan <- num
	fmt.Printf("inchant's len is %v,inchant self cap is %v \n", len(intchan), cap(intchan))
	//从管道中取出数据

	var num2 int
	num2 = <-intchan
	fmt.Println("num2=", num2)
	fmt.Printf("inchant's len is %v,inchant self cap is %v \n", len(intchan), cap(intchan))

}
