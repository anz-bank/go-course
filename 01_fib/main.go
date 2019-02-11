package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int) {
	if n == 0 {
		return
	}

	n1, n2, fibSeries := 1, 1, ""

	for i := 1; i <= abs(n); i++ {
		if n > 0 {
			//Print Fibonacci numbers for positive n
			fibSeries = fmt.Sprintf("%s%d\n", fibSeries, n1)
		} else {
			//Print Negafibonacci numbers for negative n
			fibSeries = fmt.Sprintf("%v\n%s", math.Pow(-1, float64(i+1))*float64(n1), fibSeries)
		}
		n1, n2 = n2, n1+n2
	}
	fmt.Fprint(out, fibSeries)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
