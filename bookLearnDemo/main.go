package main

import (
	"fmt"
	"reflect"
	"sort"
	"unsafe"
)

func main() {
	//数组
	//var a = [...]int{1,2,3,4}
	//var b = &a
	//fmt.Println(a[0],b[0])
	//for i,v := range b {
	//	fmt.Println(i,v)
	//}

	//var time [5][2]int
	//time[1][1] = 6
	//fmt.Println(time)
	//for range time {
	//	fmt.Println("HHH")
	//}

	//var str1 = "hello"
	//for i, _ := range str1 {
	//	fmt.Println(i, str1[:i], str1[i:])
	//}

	//fmt.Println(1<<20)

	n := Num{
		i: "aaaaa",
		j: 100,
	}
	nPoint := unsafe.Pointer(&n)
	niPoint := (*string)(unsafe.Pointer(nPoint))
	*niPoint = "kakaxi"
	njPoint := (*int64)(unsafe.Pointer(uintptr(nPoint) + unsafe.Offsetof(n.j)))
	*njPoint = 222
	fmt.Printf("n.i: %s,n.j: %d", n.i, n.j)
	fmt.Println("\n====\n")
	b1 := fn1(1, 2, 3, 4, 5, 6, 7)
	fmt.Println("b1: ", b1)

	b2 := []interface{}{123, "abc"}
	fn2(b2...)
	fn2(b2)

}

type Num struct {
	i string
	j int64
}

func SortFloat64FastV1(a []float64) {
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
	sort.Ints(b)
}

func SortFloat64FastV2(a []float64) {
	var b []int
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	bHdr := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	*bHdr = *aHdr
	sort.Ints(b)
}

func fn1(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

func fn2(a ...interface{}) {
	fmt.Println(a...)
	fmt.Print("\n")
}
