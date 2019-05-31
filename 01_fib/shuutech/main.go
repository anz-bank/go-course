package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout
var arrayList = []int{0, 1}

func nextVal(a int, b int, n int, c int) []int {
	numPair := []int{a, b}
	var i = numPair[0] + numPair[1]
	var j = numPair[1] + i
	newPair := []int{i, j}
	c += 1
	if c < n {
		arrayList = append(arrayList, newPair...)
		nextVal(i, j, n, c)
		fmt.Fprintln(out, c)

	}

	return newPair
}

func fib(n int) []int {
	var slice2 = arrayList[:n]

	return slice2
}

func main() {
	var n = 7
	nextVal(0, 1, n, 0)
	fmt.Fprintln(out, fib(n))
}
