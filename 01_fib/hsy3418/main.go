package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {

	first := 1
	second := 1
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			fmt.Fprintln(out, first)
		case 1:
			fmt.Fprintln(out, second)
		default:
			current := first + second
			first = second
			second = current
			fmt.Fprintln(out, current)
		}

	}

}

func main() {
	fib(7)
}
