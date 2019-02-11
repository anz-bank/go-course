package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fibonacciSeries(7)
}

func fibonacciSeries(number int) {
	var a, b = 0, 1
	fmt.Fprintln(out, a)
	fmt.Fprintln(out, b)
	for count := 0; count < number-2; count++ {
		var sum = a + b
		fmt.Fprintln(out, sum)
		a, b = b, sum

	}
}
