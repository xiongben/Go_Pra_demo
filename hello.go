package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(" world")
	a := buffer.String()
	fmt.Println(a)
}
