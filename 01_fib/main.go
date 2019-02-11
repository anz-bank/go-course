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

	n1 := 0
	n2 := 1

	fibSeries := ""

	for i := 0; i < n; i++ {
		if n1 == 0 {
			fibSeries = fmt.Sprintf("%s%d", fibSeries, n1)
		} else {
			fibSeries = fmt.Sprintf("%d %s %d", n1*-1, fibSeries, n1)
		}
		sum := n1 + n2
		n1 = n2
		n2 = sum
	}

	fmt.Fprint(out, fibSeries)
}
