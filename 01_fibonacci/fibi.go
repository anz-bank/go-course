package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//prints the fibonacci series
func fibonacci(n int) {
	a, b := 0, 1

	for i := 0; i <= n; i++ {
		if a >= 1 {
			fmt.Fprintln(out, a)
		}
		a, b = b, (a + b)
	}
}

func main() {
	fibonacci(7)
}
