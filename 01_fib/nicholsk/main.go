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

func fib(n int) {

	if n < 0 {
		fmt.Fprintln(out, "fib(n) doesn't accept negative integers")
		return
	}

	if n == 0 {
		fmt.Fprintln(out, n)
		return
	}

	current, previous := 1, 0

	// print first value which is always 1
	fmt.Fprintln(out, current)

	for i := 1; i < n; i++ {

		// calculate next number in sequence and print out
		sum := previous + current
		fmt.Fprintln(out, sum)

		// update variables for next iteration
		previous, current = current, sum

	}

}
