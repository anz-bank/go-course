package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// Returns Absolute vaule for a give integer
func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

// Returns Fibonacci Number of given input n (-92 <= n <= 92) - where:
// F0=0
// for n > 0
//   F1=1, Fn = Fn-1 + Fn-2
// for n < 0H
//   F-n=(-1)^(n+1)Fn.
// max n = 92
func fib(n int) {
	if abs(n) > 92 {
		fmt.Fprintln(out, `Fibonacci numbers greater than 92 not supported`)
		return
	}

	a, b := 0, 0

	for i := 0; i < abs(n)+1; i++ {
		// convert to negafibonacci number if n < 0
		if n < 0 && i%2 == 0 {
			fmt.Fprintln(out, (a+b)*-1)
		} else {
			fmt.Fprintln(out, a+b)
		}

		switch i {
		case 0:
			b = 1
		case 1:
		default:
			a, b = b, b+a
		}
	}
}

func main() {
	fib(7)
}
