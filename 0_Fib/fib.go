package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func Fibonacci(n int) []int {

	fib, t1, t2 := []int{}, 1, 1
	for i := 0; i < n; i++ {
		// add the next member
		fib = append(fib, t1)
		t1, t2 = t2, t1+t2
	}
	return fib
}

func main() {
	fmt.Fprint(out, "Fibonnaci Series of 7:", Fibonacci(7))

}
