package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

// fib computes the Fibonacci sequence
func fib(n int) {

	if n == 0 {
		fmt.Fprintln(out, n)
		return
	}
	absVal := int(math.Abs(float64(n)))
	current, previous := 1, 0

	for i := 1; i <= absVal; i++ {
		sign := 1
		if n < 0 && (i%2 == 0) { // for negative fib input, print a negative sign on even index
			sign = -1
		}
		fmt.Fprintln(out, current*sign)
		previous, current = current, (current + previous)
	}
}
func main() {
	fib(7)
}
