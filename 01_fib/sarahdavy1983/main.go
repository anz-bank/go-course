package main

import (
	"fmt"
	"io"
	"os"
)

// Write to this rather that Stdout directly to help with testing
var out io.Writer = os.Stdout

// fib returns a slice of integer fibonacci numbers
func fib(n int) []int64 {

	switch {
	case n == 0:
		return []int64{0, 0}
	case n == 1:
		return []int64{0, 1}
	case n == -1:
		return []int64{0, 1}
	case n < -1:
		x := fib(n + 1)
		nextVal := x[len(x)-2] - x[len(x)-1]
		return append(x, nextVal)
	default:
		x := fib(n - 1)
		nextVal := x[len(x)-2] + x[len(x)-1]
		return append(x, nextVal)
	}

}

//printEachindex iterates over the slice and prints each index
func printEachIndex(fibSlice []int64) {
	for i := 1; i < len(fibSlice); i++ {
		fmt.Fprintln(out, fibSlice[i])
	}
}
func main() {
	printEachIndex(fib(7))
}
