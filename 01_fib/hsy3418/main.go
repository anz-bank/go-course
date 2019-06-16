package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//fib is a function to print out a list of fibonacci number
func fib(n int) {
	first, second := 1, 1
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, first)
		first, second = second, first+second
	}
}

func main() {
	fib(7)
}
