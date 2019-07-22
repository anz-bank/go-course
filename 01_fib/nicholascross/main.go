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
	if n == 0 {
		return
	}

	prev, curr, sign := 0, 1, 1

	if n < 0 {
		sign = -1
	}

	fmt.Fprint(out, curr)

	length := sign * n
	for i := 1; i < length; i++ {
		next := prev + curr*sign
		prev = curr
		curr = next
		fmt.Fprint(out, "\n", curr)
	}
}
