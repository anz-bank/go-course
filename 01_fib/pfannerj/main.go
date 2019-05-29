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
	fib(7)
	fmt.Fprintf(mainout, "Fibonacci series completed...\n")
}

func fib(n int) {
	prev := 0
	this := 1
	nega := false //For negafibonacci - indicates whether the incoming integer value is negative
	sign := false //Used to generate the sign output for the negafibonacci series on alternate iterations
	if n < 0 {
		nega = true
		n *= -1
	}
	for i := 1; i <= n; i++ {
		if nega && sign {
			fmt.Fprintf(fibout, "-") //Print the sign for alternate iterations if nega
		}
		sign = !sign //Sign only needs to be printed on alternate iterations for negafibonacci
		fmt.Fprintln(fibout, this)
		sum := this + prev
		prev = this
		this = sum
	}
}
