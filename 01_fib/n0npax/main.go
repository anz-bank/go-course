package main

import (
	"fmt"
	"io"
	"os"
)

// Allow to overwrite stream by tests
var out io.Writer = os.Stdout

// This function is useless and it's provided just due to requirements
func main() {
	fib(7)
}

// Wrapper function to print computed Fibonacci numbers
func fib(num int) {
	for _, v := range dispatchFib(num) {
		fmt.Fprintln(out, v)
	}

}

// Dispatch and return proper numbers
func dispatchFib(num int) []int {
	switch {
	case num == 0:
		return []int{0}
	case num < 0:
		return computeNegafibonacci(num)
	default:
		return computeFibonacci(num)

	}
}

// Compute Fibonacci numbers
func computeFibonacci(num int) []int {
	fibNumbers := make([]int, num)
	a, b := 1, 1
	for i := range fibNumbers {
		fibNumbers[i] = a
		a, b = b, a+b
	}
	return fibNumbers
}

// Compute negafibonacci numbers
func computeNegafibonacci(num int) []int {
	fibNumbers := computeFibonacci(num * -1)
	for i, v := range fibNumbers {
		fibNumbers[i] = v * [2]int{1, -1}[i%2]
	}
	return fibNumbers
}
