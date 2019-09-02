package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(i int) {
	m, n := 0, 1
	if i >= 0 {
		for j := 0; j < i; j++ {
			m += n
			m, n = n, m
			fmt.Fprintln(out, m)
		}
	} else {
		for j := 0; j > i; j-- {
			m -= n
			m, n = n, m
			fmt.Fprintln(out, m)
		}
	}
}

// main function
func main() {
	fib(7)
}
