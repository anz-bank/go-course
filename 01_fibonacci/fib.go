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
	prev, cur := 0, 1
	var fibSeries string

	for i := 1; i <= n; i++ {
		sum := 0
		if i == n {
			fibSeries += strconv.Itoa(prev)
		} else {
			fibSeries += strconv.Itoa(prev) + " "
		}
		sum = prev + cur
		prev = cur
		cur = sum
	}
	fmt.Fprintln(out, fibSeries)
}
