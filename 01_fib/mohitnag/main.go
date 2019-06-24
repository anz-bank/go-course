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

func fibonacci(n int) []int {
	if n < 1 {
		return []int{}
	}
	slice := make([]int, n+1)
	slice[0], slice[1] = 1, 1
	for i := 2; i < n; i++ {
		slice[i] = slice[i-1] + slice[i-2]
	}
	return slice[:n]
}

func fib(n int) {
	slice := fibonacci(n)
	for i := range slice {
		fmt.Fprintln(out, slice[i])
	}
}
