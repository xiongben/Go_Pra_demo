package main

func addUpper(n int) int {
	res := 0
	for i := 0; i <= n-1; i++ {
		res += i
	}
	return res
}

func getSub(n1, n2 int) int {
	return n1 - n2
}
