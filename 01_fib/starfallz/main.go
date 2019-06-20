package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	printIntSlice(fib(7))
}

func printIntSlice(input []int) {
	for _, num := range input {
		fmt.Fprintln(out, num)
	}
}

func fib(n int) []int {
	sign := 1
	if n < 0 {
		sign = -1
	}

	n1, n2 := 0, 1
	result := make([]int, n*sign)
	for i := 0; i < n*sign; i++ {
		n1, n2 = n2, n1+sign*n2
		result[i] = n1
	}
	return result
}
