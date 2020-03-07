package main

import (
	"fmt"
	"testing"
)

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 45 {
		//fmt.Printf("addUpper(10）执行错误，期望值是%v,实际值是%v \n",55,res)
		t.Fatalf("addUpper(10）执行错误，期望值是%v,实际值是%v \n", 55, res)
	}
	t.Logf("addupper(10)执行正确，，，")
}

func TestGetSub(t *testing.T) {
	res := getSub(10, 5)
	if res != 5 {
		//fmt.Printf("addUpper(10）执行错误，期望值是%v,实际值是%v \n",55,res)
		t.Fatalf("getSub(10,5）执行错误，期望值是%v,实际值是%v \n", 5, res)
	}
	t.Logf("getSub(10,5)执行正确，，，")
}

func TestHello(t *testing.T) {
	fmt.Println("this is hello test")
}
