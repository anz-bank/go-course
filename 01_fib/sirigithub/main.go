package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//function returns absolute value of an interger
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// fibonaci series
func fib(n int) {

	if n == 0 {
		fmt.Fprintln(out, n)
		return
	}
	absVal := abs(n)
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
