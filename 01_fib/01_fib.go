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
func fib(n int) int {
	var num1 = 0
	var num2 = 1
	var sum = 0
	fmt.Fprintln(out, num2)
	for i := 1; i < n; i++ {
		sum = num1 + num2
		num1 = num2
		num2 = sum
		fmt.Fprintln(out, sum)
	}

	return sum
}
