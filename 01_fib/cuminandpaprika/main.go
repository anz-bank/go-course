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

// Print n number of fibonnaci sequence
// 1, 1, 2, 3, 5, 8, 13
// Note: we take the absolute value of n to ensure our loop works.
func fib(n int64) {
	var i int64
	if isPositive(n) {
		for ; i < n; i++ {
			fmt.Fprintln(out, calcPositiveFib(i))
		}
		return
	}
	if n == 0 {
		fmt.Fprintln(out, calcZeroFib(i))
	}

	for ; i > n; i-- {
		fmt.Fprintln(out, calcNegativeFib(i))
	}

}

func calcZeroFib(n int64) int64 {
	return 0
}

func calcPositiveFib(n int64) int64 {
	if n == 0 {
		return calcZeroFib(0)
	}
	if n == 1 {
		return 1
	}

	return calcPositiveFib(n-2) + calcPositiveFib(n-1)
}

func calcNegativeFib(n int64) int64 {
	// We convert the negative to positive so it calculates correctly with calcPositiveFib
	n = (-1 * n)
	if isEven(n) {
		return ((-1) * calcPositiveFib(n))
	}
	return calcPositiveFib(n)

}

func isEven(n int64) bool {
	return ((n & 1) == 0)
}

func isPositive(n int64) bool {
	return (n > 0)
}
