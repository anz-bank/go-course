package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	if n < 1 || n > 92 {
		fmt.Fprintln(out, "fib.go only handles integers between 1 to 92 inclusive")
		return
	}

	for _, num := range genFib(n) {
		fmt.Fprintln(out, num)
	}
}

func genFib(n int) []int64 {
	series := make([]int64, n)
	series[0] = 1

	if n == 1 {
		return series
	}

	series[1] = 1
	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}

	return series
}

func main() {
	fib(7)
}
