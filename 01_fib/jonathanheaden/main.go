package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int) {
	a, b := 0, 1
	var prefix string
	var neg = false
	if n < 0 {
		n *= -1
		neg = true
	}
	for i := 0; i <= n; i++ {
		if prefix = ""; neg && (i%2 == 0) {
			prefix = "-"
		}
		switch i {
		case 0:
			if n == 0 {
				fmt.Fprintln(out, "0")
			}
		case 1:
			fmt.Fprintf(out, "%s%d\n", prefix, 1)
		default:
			a, b = b, a+b
			fmt.Fprintf(out, "%s%d\n", prefix, b)
		}
	}
}
