package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fibonacciSeries(7)
}

func fibonacciSeries(number int) {
	if number == 0 {
		return
	}
	var a, b = 1, 1
	for count := 0; count < abs(number); count++ {
		if number > 0 {
			fmt.Fprintln(out, a)
		} else {
			fmt.Fprintln(out, math.Pow(-1, float64(count))*float64(a))
		}
		a, b = b, a+b
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
