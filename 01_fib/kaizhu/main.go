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

func fib(n int64) {
	switch {
	// print 0 if n is 0
	case n == 0:
		fmt.Fprintln(out, 0)
	// print a message if n is too big or too small
	case n < -92 || n > 92:
		fmt.Fprintln(out, "Please enter a number within the range [-92, 92].")
	// print fib array if n is within range
	default:
		// check if n is negative
		var sign int64 = 1
		if n < 0 {
			n = 0 - n
			sign = -1
		}

		// generate fib array
		// fibArray := make([]int64, n)
		var a, b int64 = 1, sign
		var i int64
		for i = 0; i < n; i++ {
			fmt.Fprintln(out, a)
			a, b = b, a+b*sign
		}
	}
}
