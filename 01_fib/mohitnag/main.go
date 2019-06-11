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
	slice := []int{}
	if n < 1 {
		return []int{}
	}
	if n == 1 {
		return []int{1}
	}
	slice = append(slice, 1, 1)
	for i := 2; i < n; i++ {
		slice = append(slice, slice[i-1]+slice[i-2])
	}
	return slice
}

func fib(n int) {
	slice := fibonacci(n)
	for i := range slice {
		fmt.Fprintln(out, slice[i])
	}
}
