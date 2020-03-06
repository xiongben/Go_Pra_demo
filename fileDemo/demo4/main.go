package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filepath := "/Users/mozat/go/src/Go_Pra_demo/fileDemo/b.txt"
	//追加内容
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
		return
	}
	defer file.Close()

	//读取显示在终端
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}

	str := "add somthing in file now ~~~ \r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
