package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func abs(args int) (res int) {
	if args == 0 {
		return 0
	}
	if args > 0 {
		return args
	}
	return args * (-1)

}

func fib(args int) {
	a, b := 0, 1
	if abs(args) > 92 {
		fmt.Fprint(out, "Fibonacci numbers greater than 92 not supported")
		return
	}
	if abs(args) <= 1 {
		fmt.Fprintln(out, args)
	} else {
		for index := 1; index <= abs(args); index++ {
			if index < 2 {
				fmt.Fprintln(out, b)
			} else {
				res := a + b
				a, b = b, res
				if args < 0 && index%2 == 0 {
					fmt.Fprintln(out, res*(-1))
				} else {
					fmt.Fprintln(out, res)
				}
			}
		}
	}
}

func main() {
	fib(7)
}
