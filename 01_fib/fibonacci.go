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

func fib(num int) {
	num1, num2 := 1, 1
	for i := 0; i < num; i++ {
		fmt.Fprintln(out, num1)
		num1, num2 = num2, num1+num2
	}
}
