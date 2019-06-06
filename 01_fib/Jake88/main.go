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
		n = n * -1
		negative = true
	}
	return n, negative
}

func fib(n int) {
	// Cator for fib(0)
	if n == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	prev2, prev1 := 1, 0

	n, nega := absolute(n)

	for i := 0; i < n; i++ {
		current := prev1 + prev2

		}

		prev2, prev1 = prev1, current

		// print based on whether we are in negafibonacci or not
		if nega && i%2 == 1 {
			fmt.Fprintln(out, -current)
		} else {
			fmt.Fprintln(out, current)
		}
	}
}
