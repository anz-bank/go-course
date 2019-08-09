package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	_ = fib(7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func fibonacciSign(n int) int {
	if n < 0 {
		return [2]int{1, -1}[abs(n+1)%2]
	}
	return 1
}

func fib(n int) error {
	// I have picked an arbitrary limit because this
	// is a learning exercise
	if n > 60 || n < -60 {
		return errors.New("fib only works between 60 and -60")
	}

	if n >= 0 {
		//  iterate up
		for idx := 0; idx <= n; idx++ {
			fmt.Fprintln(out, fibonacci(idx))
		}
	} else {
		// iterate down
		for idx := 0; idx >= n; idx-- {
			fmt.Fprintln(out, fibonacci(idx))
		}
	}
	return nil
}

func fibonacci(n int) int {
	return fibonacciSign(n) * fibonacciLoop(abs(n))
}

// could optimise this by passing in the last fibonacci
// number for this particular example but for a general
// case lets not assume we have context

func fibonacciLoop(n int) int {
	if n < 2 {
		return n
	}
	result := 0
	fibNum1 := 0
	fibNum2 := 1
	for idx := 2; idx <= n; idx++ {
		result = fibNum1 + fibNum2
		fibNum1 = fibNum2
		fibNum2 = result
	}
	return result
}
