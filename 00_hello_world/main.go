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
	for count := 0; count < number; count++ {
		fmt.Fprintln(out, a)
		a, b = b, a+b
	}
}
