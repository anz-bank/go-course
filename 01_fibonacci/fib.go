package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fibonacci(10)
}

func fibonacci(n int) {

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
