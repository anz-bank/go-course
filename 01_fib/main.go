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

func fib(n int) {

	n1, n2 := 1, 1

	fibSeries := "0"

	for i := 1; i < n; i++ {

		fibSeries = fmt.Sprintf("%d %s %d", -n1, fibSeries, n1)

		sum := n1 + n2
		n1, n2 = n2, sum
	}

	fmt.Fprint(out, fibSeries)
}
