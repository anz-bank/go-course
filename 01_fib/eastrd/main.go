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
	prev_prev, prev := 1, 1
	fmt.Fprintln(out, prev_prev)
	if n == 1 {
		return
	}
	fmt.Fprintln(out, prev)
	if n == 2 {
		return
	}

	for i := 2; i < n; i++ {
		prev_prev, prev = prev, prev_prev+prev
		fmt.Fprintln(out, prev)
	}
}

func main() {
	// fib(7)
}
