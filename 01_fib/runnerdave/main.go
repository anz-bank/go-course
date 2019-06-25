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
		return
	}
	// 1 for positive n ; -1 for negative n
	sign := 1
	if n < 0 {
		sign = -1
	}
	n1, n2 := 1, sign
	for j := 0; j < n*sign; j++ {
		fmt.Fprintln(out, n1)
		n1, n2 = n2, n1+sign*n2
	}
}

func main() {
	fib(7)
}
