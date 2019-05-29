package main

import (
	"fmt"
	"io"
	"os"
)

var cache = make(map[int]int)
var out io.Writer = os.Stdout

func fib(n int) int {
	if cacheVal, ok := cache[n]; ok {
		return cacheVal
	}
	switch {
	case n == 1, n == 0:
		//print 1 once
		if _, ok := cache[1]; !ok {
			fmt.Fprintln(out, "1")
		}
		cache[n] = n
		return n
	case n < 0:
		sign := 1
		if n%2 == 0 {
			sign = -1
		}
		nval := sign * fib(-n)
		cache[n] = nval
		fmt.Fprintln(out, nval)
		return nval
	default:
		val := fib(n-1) + fib(n-2)
		cache[n] = val
		fmt.Fprintln(out, val)
		return val
	}
}

func main() {
	fib(7)
}
