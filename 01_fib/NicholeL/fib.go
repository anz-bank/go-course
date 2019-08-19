package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * (-1)
}

func fib(arg int) {
	a, b := 0, 1
	if abs(arg) <= 1 {
		fmt.Fprintln(out, arg)
		return
	}
	for index := 1; index <= abs(arg); index++ {
		if index < 2 {
			fmt.Fprintln(out, b)
		} else {
			res := a + b
			a, b = b, res
			if arg < 0 && index%2 == 0 {
				fmt.Fprintln(out, res*(-1))
			} else {
				fmt.Fprintln(out, res)
			}
		}
	}
}

func main() {
	fib(7)
}
