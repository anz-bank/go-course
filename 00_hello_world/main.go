package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(i int) int {
	a, b := 0, 1
	fmt.Fprintln(out, b)
	sum := 0
	for cnt := 0; cnt < i-1; cnt++ {
		sum = a + b
		fmt.Fprintln(out, sum)
		a, b = b, sum
	}
	return sum
}

func main() {
	fmt.Fprintln(out, "Hallo du schöne Welt!")
	fib(7)
}
