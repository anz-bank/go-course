package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	if n == 0 {
		fmt.Fprintln(out, 0)
	}

	a, b := 1, 1
	for i := 0; i < abs(n); i++ {
		if n < 0 {
			fmt.Fprintln(out, coefficient(uint(i+1))*a)
		} else {
			fmt.Fprintln(out, a)
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

func coefficient(n uint) int {
	if (n+1)%2 == 0 {
		return 1
	}
	return -1
}

func main() {
	fib(7)
}
