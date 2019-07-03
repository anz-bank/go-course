package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Returns Fibonacci and Negafibonacci series
func fib(n int) {
	m := Abs(n)
	a, b := 0, 1
	for i := 1; i <= m; i++ {
		sign := 1
		if (i%2) == 0 && n < 0 {
			sign = -1
		}
		fmt.Fprintln(out, sign*b)
		a, b = b, a+b
	}
}

func main() {
	fib(7)
	fib(-7)
}
