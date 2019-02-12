package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	var num1, num2 = 1, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, num1)
		num1, num2 = num2, num1+num2
	}
}

func main() {
	fib(7)
}
