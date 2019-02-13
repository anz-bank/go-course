package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, a)
		a, b = b, a+b
	}
}

func main() {
	fib(7)
}
