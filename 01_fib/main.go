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

	if n <= 0 {
		return
	}

	n1, n2, fibSeries := 1, 1, "0\n"

	for i := 1; i < n; i++ {
		fibSeries = fmt.Sprintf("%d\n%s%d\n", -n1, fibSeries, n1)
		n1, n2 = n2, n1+n2
	}

	fmt.Fprint(out, fibSeries)
}
