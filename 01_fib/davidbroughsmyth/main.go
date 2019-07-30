package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// Function to calculate Fibonacci or Negafibonacci number
// Range of n from -92 .. 0 .. 92
func fib(n int) {

	if n == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	if n > 92 || n < -92 {
		fmt.Fprintf(out, "Range for fib(n) should be between -92 to 92\n")
		return
	}

	var n2, n1, num int = 0, 1, n
	if n < 0 {
		num *= -1
	}

	for i := 0; i < num; i++ {
		if n < 0 && i%2 != 0 {
			fmt.Fprintln(out, n1*-1)
		} else {
			fmt.Fprintln(out, n1)
		}
		n2, n1 = n1, n1+n2
	}
}

func main() {

	fib(7)

}
