package main

import (
	"fmt"
)

var arrayList = []int{0, 1}

func nextVal(a int, b int, n int, c int) []int {
	numPair := []int{a, b}
	var i = numPair[0] + numPair[1]
	var j = numPair[1] + i
	newPair := []int{i, j}
	c++
	if c < n {
		arrayList = append(arrayList, newPair...)
		nextVal(i, j, n, c)

	}

	return newPair
}

func fib(n int) []int {
	nextVal(0, 1, n, 0)
	var slice2 = arrayList[:n]
	fmt.Println(slice2)
	return slice2
}

func main() {
	fib(7)
}
