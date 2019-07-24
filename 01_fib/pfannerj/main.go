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

func fib(n int) { //NB: This caters for int values in the range -92 to 92
	prev := 0
	this := 1
	nega := false //For negafibonacci - indicates whether the incoming integer value is negative
	if n < 0 {
		nega = true
		n *= -1
	}
	if n > 92 {
		fmt.Fprintf(fibout, "Value outside allowable range (-92 to 92)\n")
		return
	}
	for i := 1; i <= n; i++ {
		if nega && (i%2 == 0) { //Negative sign needs to be printed on alternate iterations for negafibonacci
			fmt.Fprintf(fibout, "-") //Print the negitive sign
		}
		fmt.Fprintln(fibout, this)
		prev, this = this, this+prev
	}
}
