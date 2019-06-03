package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fmt.Fprintln(out, "Call Fibonacci numbers!")
	fib(7)
}

func fib(n int) {

	current, next := 1, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, current)
		current, next = next, current+next
	}
}
