package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

// var (
// 	sum       int
// 	prev      int    = 0
// 	cur       int    = 1
// 	fibSeries string = ""
// )

func main() {
	fibonacci(10)
}

func fibonacci(n int) {
	// var sum int
	// var prev int = 0
	// var cur int = 1
	// var fibSeries string = ""

	var sum int
	var prev, cur int = 0, 1
	var fibSeries string

	for i := 1; i <= n; i++ {
		// fmt.Println(prev)
		fibSeries += strconv.Itoa(prev) + " "
		sum = prev + cur
		prev = cur
		cur = sum
	}
	fmt.Fprintln(out, fibSeries)
}
