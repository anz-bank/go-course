package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, b)
		a, b = b, a+b
	}
}

func main() {
	fib(7)
}
