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
	isPositive := true
	switch {
	case n == 0:
		fmt.Fprintln(out, 0)
	case n < 0:
		isPositive = false
	}

	first, count := 1, 1
	second, sum := 0, 0
	n = int(math.Abs(float64(n)))
	for count <= n {
		sum = first + second
		if !isPositive && count%2 == 0 {
			fmt.Fprintln(out, sum*-1)
		} else {
			fmt.Fprintln(out, sum)
		}
		first, second = second, sum
		count++
	}
}
