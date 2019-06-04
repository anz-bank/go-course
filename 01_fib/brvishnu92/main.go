package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int) {
	isPos := true
	switch {
	case n == 0:
		fmt.Fprintln(out, 0)
	case n < 0:
		isPos = false
	}

	n1 := 1
	n2 := 0
	n = int(math.Abs(float64(n)))
	for i := 1; i <= n; i++ {
		n1, n2 = n2, n1+n2
		if !isPos && i%2 == 0 {
			fmt.Fprintln(out, n2*-1)
		} else {
			fmt.Fprintln(out, n2)
		}
	}
}
