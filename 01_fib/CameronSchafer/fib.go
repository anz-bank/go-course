package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int) {
	sign, n1, n2 := 1, 0, 1
	//check if use nega or normal fib sequence
	if n < 0 {
		sign = -1
	}

	n *= sign
	//loop until all numbers have been printed
	for i := 0; i < n; i++ {
		if i%2 == 0 || sign > 0 {
			//will get the negafib number if the sign is -1
			n1, n2 = n2, n1+n2*sign
		} else {
			//will account for the negative number during the negafib sequence
			n1, n2 = n2, n1-n2
		}
		fmt.Fprintln(out, n1)
	}
}

func main() {
	fib(7)
}
