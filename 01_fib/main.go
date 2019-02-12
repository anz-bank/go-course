package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(i int) {
	a, b := 1, 1
	for cnt := 0; cnt < i; cnt++ {
		fmt.Fprintln(out, a)
		a, b = b, a+b
	}
}

func main() {
	fib(7)
}
