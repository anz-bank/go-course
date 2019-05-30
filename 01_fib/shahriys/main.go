package main

import (
	"fmt"
	"io"
	"os"
)

var outfib io.Writer = os.Stdout

// Function to calculate nth Fibonacci or Negafibonacci number
func nfib(n int) int64 {

	if n >= 0 {

		switch n {
		case 0, 1:
			return 1
		default:
			return nfib(n-1) + nfib(n-2)
		}
	} else {

		switch n {
		case -1:
			return 1
		case -2:
			return -1
		default:
			return nfib(n+2) - nfib(n+1)
		}
	}
}

// Function to printh first n Fibonacci or Negafibonacci numbers
func fib(n int) {
	if n >= 0 {
		for i := 0; i < n; i++ {
			fmt.Fprintln(outfib, nfib(i))
		}
	} else {
		for i := -1; i >= n; i-- {
			fmt.Fprintln(outfib, nfib(i))
		}
	}
}

func main() {
	fib(7)
	fib(-7)
}
