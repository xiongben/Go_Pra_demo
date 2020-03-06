package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("/Users/mozat/go/src/Go_Pra_demo/fileDemo/a.txt")
	if err != nil {
		fmt.Println("open file error=", err)
	}
	fmt.Println(file)
	//err = file.Close()
	//if err != nil {
	//	fmt.Println("close file err=",err)
	//}

	//打开方式一，带缓存的方式
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		fmt.Print(str)
		if err == io.EOF {
			break
		}

	}
	fmt.Println("文件读取结束")

	//打开方式二，用于文件较小的情况
	filepath := "/Users/mozat/go/src/Go_Pra_demo/fileDemo/a.txt"
	content1, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("read file err=", err)
	}
	fmt.Println(content1) //byte[]
}
