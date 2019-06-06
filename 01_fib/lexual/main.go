package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

// fib prints out n fibonacci numbers, use negafibonacci if n < 0
func fib(n int) {
	isPos := n > 0
	if !isPos {
		n *= -1
	}
	a, b := 0, 1
	var f int
	for i := 1; i <= n; i++ {
		if isPos {
			f = b
		} else {
			// negafibonnaci: fib(-n) = (-1)**(n+1) * fib(n)
			f = int(math.Pow(-1, float64(i+1))) * b
		}
		fmt.Fprintln(out, f)
		a, b = b, a+b
	}
}

func main() {
	fib(7)
}
