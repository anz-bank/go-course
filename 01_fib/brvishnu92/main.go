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

	n1, count := 1, 1
	n2, sum := 0, 0
	n = int(math.Abs(float64(n)))
	for count <= n {
		sum = n1 + n2
		if !isPos && count%2 == 0 {
			fmt.Fprintln(out, sum*-1)
		} else {
			fmt.Fprintln(out, sum)
		}
		n1, n2 = n2, sum
		count++
	}
}
