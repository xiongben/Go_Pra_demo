package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filepath := "/Users/mozat/go/src/Go_Pra_demo/fileDemo/b.txt"
	//追加内容
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()
	str := "add somthing in file now \r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
