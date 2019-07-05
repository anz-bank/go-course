package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	if n <= 0 {
		fmt.Fprintln(out, "Invalid Input: Must be a positive integer")
		return
	}

	current, previous := 1, 0
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, current)
		previous, current = current, current+previous
	}
}

func main() {
	fib(7)
}
