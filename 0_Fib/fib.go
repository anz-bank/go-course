package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func Fibonacci(n int) []int {

	fib := []int{}
	t1 := 0
	t2 := 1
	sum := 0
	for i := 0; i <= n; i++ {
		// add the next member
		fib = append(fib, t1)
		sum = t1 + t2
		t1 = t2
		t2 = sum
	}

	return fib
}

func main() {
	fmt.Fprint(out, "Fibonnaci Series of 8:", Fibonacci(8))

}
