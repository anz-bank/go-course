package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Write to this rather that Stdout directly to help with testing.
var out io.Writer = os.Stdout

// `fibSlice` returns a slice of integer fibonacci numbers starting at 0.
func fibSlice(n int) []int64 {
	switch {
	case n > 92:
		panic("Fibonnaci value is to large for int64")
	case n < -92:
		panic("Fibonnaci value is to large for int64")
	case n == 0:
		return []int64{0}
	case n == 1 || n == -1:
		return []int64{0, 1}
	case n < -1:
		a := fibSlice(n + 1)
		nextVal := a[len(a)-2] - a[len(a)-1]
		return append(a, nextVal)
	default:
		a := fibSlice(n - 1)
		nextVal := a[len(a)-2] + a[len(a)-1]
		return append(a, nextVal)
	}
}

// `fibString` returns a newline delimited string of fibonacci numbers starting at 0.
// A trailing newline is included.
func fibString(n int) string {
	var b strings.Builder
	for _, v := range fibSlice(n) {
		fmt.Fprintln(&b, v)
	}
	return b.String()
}

// `fib` prints a newline dxelimited string of fibonacci numbers starting at 0.
// A trailing newline is included.
func fib(n int) {
	fmt.Fprint(out, fibString(n))
}

func main() {
	fmt.Fprint(out, "\nFib 7\n------\n")
	fib(7)
}
