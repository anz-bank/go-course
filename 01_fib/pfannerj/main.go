package main

import (
	"fmt"
	"io"
	"os"
)

var fibout io.Writer = os.Stdout
var mainout io.Writer = os.Stdout

func main() {
	fmt.Fprintf(mainout, "Fibonacci series starting...\n")
	f := 7
	fib(f)
	fmt.Fprintf(mainout, "Fibonacci series completed...\n")
}

func fib(n int) {
	prev := 0
	this := 1
	nega := false
	sign := false
	if n < 0 {
		nega = true
		n *= -1
	}
	for i := 1; i <= n; i++ {
		if nega && sign {
			fmt.Fprintf(fibout, "-")
		}
		sign = !sign
		fmt.Fprintln(fibout, this)
		sum := this + prev
		prev = this
		this = sum
	}
}
