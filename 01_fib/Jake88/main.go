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
	// Our initial map, with the default starting values needed to begin the sequence.
	fibonacciMap := map[int]int{
		-1: 1,
		0:  0,
	}
	negative := false

	// Check if we are performaing negafibonacci or not, and if so remove the sign from n
	if n < 0 {
		n *= -1
		negative = true
	}

	// Form our map of fibonacci numbers, ensuring i matches Fn
	for i := 1; i < n+1; i++ {
		// Store the next positive fibonacci iteration
		current := fibonacciMap[i-1] + fibonacciMap[i-2]
		fibonacciMap[i] = current

		// If negafibonacci, store the next negative iteration
		if negative {
			if i%2 == 0 {
				current = -current
			}
			fibonacciMap[-i] = current
		}

		// Print out the current iteration value
		fmt.Fprintln(out, current)
	}
}
