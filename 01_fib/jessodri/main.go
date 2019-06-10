package main

import (
	"fmt"
	"io"
	"math/big"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {

	a, b, c := big.NewInt(0), big.NewInt(1), 1
	if n < 0 {
		c = -1
	}
	for i := 1; i <= n*c; i++ {
		isNeg, fibResult := big.NewInt(1), big.NewInt(0)

		if i%2 == 0 && n < 0 {
			isNeg = big.NewInt(-1)
		}
		a, b = b, a.Add(a, b)
		fibResult.Mul(a, isNeg)

		fmt.Fprintln(out, fibResult)
	}
}

func main() {
	fib(7)
}
