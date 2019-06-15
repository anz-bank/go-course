package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fib(7) //Provide the input, programmed for both Positive and Negative Series
}

func abs(x int) int {

	if x < 0 {
		return -x
	}
	return x
}

func fib(n int) {

	var a int //first term
	var b = 1 //second term
	var c int //next term
	var absn = abs(n)
	var sign = n / absn

	for i := 0; i < absn; i++ {
		if i > 0 {
			c = a + b*sign
			a = b
			b = c
		} else {
			c = a + b
		}
		fmt.Fprintln(out, c)
	}
}
