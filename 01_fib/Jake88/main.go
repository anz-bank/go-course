package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func absolute(n int) (int, bool) {
	var negative bool
	if n < 0 {
		n *= -1
		negative = true
	}
	return n, negative
}

func fib(n int) {
	// Cater for fib(0).
	if n == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	prev2, prev1 := 1, 0

	n, nega := absolute(n)

	for i := 0; i < n; i++ {
		current := prev1 + prev2

		// Protect against integer overflow / out of bounds error.
		if current < prev1 {
			fmt.Fprintln(out, "Error: Number overflow")
			return
		}

		prev2, prev1 = prev1, current

		// print based on whether negafibonacci was requested or not.
		if nega && i%2 == 1 {
			fmt.Fprintln(out, -current)
		} else {
			fmt.Fprintln(out, current)
		}
	}
}
