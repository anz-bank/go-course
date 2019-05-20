package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fib(n int) {
	if n == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	absVal := abs(n)
	sign := n / absVal
	n1, n2 := 1, sign

	for j := 0; j < absVal; j++ {
		fmt.Fprintln(out, n1)
		n1, n2 = n2, n1+sign*n2
	}
}

func main() {
	fib(7)
}
