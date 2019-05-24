package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) int {
	absoluteValue := int(math.Abs(float64(n)))
	f := make([]int, absoluteValue+1, absoluteValue+2)
	f[0] = 0
	switch {
	case n < 0:
		f[1] = -1
	case n > 0:
		f[1] = 1
	default:
		return f[0]
	}
	fmt.Fprintf(out, "%d\n", f[1])
	for i := 2; i <= absoluteValue; i++ {
		f[i] = f[i-1] + f[i-2]
		fmt.Fprintf(out, "%d\n", f[i])
	}
	return f[absoluteValue]
}

func main() {
	fib(7)
}
