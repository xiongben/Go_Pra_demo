package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type charCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	fileName := "/Users/xiongben/go/src/awesomeProject1/fileDemo/a.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	var count charCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//遍历str，进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("char number is %v ,num number is %v,space number is %v,other number is %v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
