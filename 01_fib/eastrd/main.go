package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	if n <= 0 {
		fmt.Fprintln(out, "Fibonacci number has to be positive")
		return
	}
	prevPrev, prev := 1, 1
	fmt.Fprintln(out, prevPrev)
	if n == 1 {
		return
	}
	fmt.Fprintln(out, prev)
	if n == 2 {
		return
	}

	for i := 2; i < n; i++ {
		prevPrev, prev = prev, prevPrev+prev
		fmt.Fprintln(out, prev)
	}
}

func main() {
	// fib(7)
}
