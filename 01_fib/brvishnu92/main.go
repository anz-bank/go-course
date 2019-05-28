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

	switch {
	case n > 0:
		printFib(n, "positive")
	case n < 0:
		printFib(n, "negative")
	default:
		fmt.Fprintln(out, 0)
	}
}

func printFib(n int, valtype string) {

	first := 1
	second := 0
	count := 1
	sum := 0
	n = int(math.Abs(float64(n)))
	for count <= n {
		sum = first + second
		if valtype == "negative" && count%2 == 0 {
			fmt.Fprintln(out, sum*-1)
		} else {
			fmt.Fprintln(out, sum)
		}
		first = second
		second = sum
		count++
	}
}
