package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fmt.Fprintln(out, "Call Fibonacci numbers!")
	// var previous = 0
	// var current = 1
	// var next = current + previous

	fib(7)
}

func fib(n int) {
	i := 0
	previous := 0
	current := 1
	next := current + previous
	for i < n {
		next = current + previous
		previous = current
		current = next
		//fmt.Println("i", i, "n", n, "previous", previous, "current", current, "next", next)
		fmt.Fprintln(out, previous)
		i++
	}
	//fmt.Println("sum 2:", next)

	// if n > 0 {
	// 	fib(n - 1)

	// 	fmt.Fprintln(out, next)
	// 	next = current + previous
	// 	previous = current
	// 	current = next
	// }

}
