package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fmt.Fprintln(out, "Call Fibonacci numbers!")

	// fmt.Println(1)
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(5)
	// fmt.Println(8)
	// fmt.Println(13)
	fib(7)
}

var previous = 0
var current = 1
var next = current + previous

func fib(n int) {

	if n > 0 {
		fib(n - 1)
		//fmt.Println("n: ", n, "previous: ", previous, ", current: ", current,", next", next)

		fmt.Fprintln(out, next)
		next = current + previous
		previous = current
		current = next
	}

}
