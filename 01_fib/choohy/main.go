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
	i := 0
	previous := 0
	current := 1
	next := 1
	for i < n {
		next = current + previous
		previous = current
		current = next
		fmt.Fprintln(out, previous)
		i++
	}
}
