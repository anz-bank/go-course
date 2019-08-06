package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	if n >= 93 || n <= -93 {
		fmt.Fprintln(out, "fib: Range out of bounds; takes input > -93 and < 93")
		return
	}

	sign := 1
	if n < 0 {
		sign = -1
	}
	sum := 0
	a := 0
	b := 1
	for i := 0; i < n*sign; i++ {
		sum = a + b*sign
		a = b
		b = sum
		fmt.Fprintln(out, a)
	}
}

func main() {
	fib(7)
}
