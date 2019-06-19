package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fibSeries(n int) []int {
	if n >= 93 || n <= -93 {
		panic("Numbers beyond 93 or -93 are not supported!")
	}

	sign := 1
	if n < 0 {
		sign = -1
	}

	series := make([]int, n*sign+1)
	if n == 0 {
		series[0] = 0
	} else {
		series[0], series[1] = 0, 1
		for i := 2; i <= n*sign; i++ {
			series[i] = series[i-2] + series[i-1]*sign
		}
	}
	return series
}

func fib(n int) {
	series := fibSeries(n)
	for _, fn := range series {
		fmt.Fprintln(out, fn)
	}
}

func main() {
	fib(7)
}
