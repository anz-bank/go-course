package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int64) []int64 {
	if n <= -93 || n >= 93 {
		fmt.Fprintln(out, "fib: Range out of bounds; takes input > -93 and < 93")
		return nil
	}

	var sign int64 = 1
	if n < 0 {
		sign = -1
	}

	var a, b, i int64
	a, b = 0, 1
	series := make([]int64, 0, n*sign)
	for i = 0; i < n*sign; i++ {
		a, b = b, a+sign*b
		series = append(series, a)
	}
	return series
}

func main() {
	fmt.Fprintln(out, fib(7))
}
