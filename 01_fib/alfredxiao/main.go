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
	var series []int
	fa, fb := 0, 1
	p, fn := 0, 0
	for {
		switch {
		case p == 0:
			fn = fa
		case p == 1:
			fn = fb
		case p > 1:
			fa, fb = fb, fa+fb
			fn = fb
		case p < 0:
			fa, fb = fb-fa, fa
			fn = fa
		}
		series = append(series, fn)
		if p == n {
			break
		}
		if n > 0 {
			p++
		} else if n < 0 {
			p--
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
