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
func fib(n int16) {
	var i int16
	if isPositive(n) {
		for ; i < n; i++ {
			fmt.Fprintln(out, calcFib(i))
		}
		return
	}

	for ; i > n; i-- {
		fmt.Fprintln(out, calcFib(i))
	}

}

func calcFib(n int16) int16 {
	if isPositive(n) {
		return calcPositiveFib(n)
	}
	return calcNegativeFib(n)
}

func calcPositiveFib(n int16) int16 {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return calcFib(n-2) + calcFib(n-1)
}

func calcNegativeFib(n int16) int16 {
	// We convert the negative to positive so it calculates correctly with calcPositiveFib
	n = (-1 * n)
	if isEven(n) {
		return ((-1) * calcPositiveFib(n))
	}
	return calcPositiveFib(n)

}

func isEven(n int16) bool {
	return ((n % 2) == 0)
}

func isPositive(n int16) bool {
	return (n >= 0)
}
