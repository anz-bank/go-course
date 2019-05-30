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
	prev, curr := 1, 1
	fmt.Fprintln(out, prev)
	if n == 1 {
		return
	}
	fmt.Fprintln(out, curr)
	if n == 2 {
		return
	}

	for i := 2; i < n; i++ {
		prev, curr = curr, prev+curr
		fmt.Fprintln(out, curr)
	}
}

func main() {
	fib(3)
}
