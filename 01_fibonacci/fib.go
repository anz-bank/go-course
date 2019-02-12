package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fibonacci(7)
}

func fibonacci(n int) {
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, a)
		a, b = b, a+b
	}
}
