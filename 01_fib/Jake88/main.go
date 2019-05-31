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
	store := map[int]int{
		0: 0,
	}

	i := 0
	for i != n {
		current := 0

		if n < 0 {
			current = store[i] - store[i-1]
			if current == 0 {
				current = 1
			}
			store[i-2] = current
			i--
		} else {
			current = store[i] + store[i-1]
			if current == 0 {
				current = 1
			}
			store[i+1] = current
			i++
		}
		fmt.Fprintln(out, current)
	}
}
