package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(n int) {
	if n < 0 {
		log.SetFlags(0)
		log.Printf("Fib of %d is not available. Defaulting to minimum fib of 0", n)
		n = 0
	} else if n > 45 {
		log.SetFlags(0)
		log.Printf("Fib of %d is not available. Defaulting to maximum fib of 45", n)
		n = 45
	}
	fibSlice := []int{}
	nextFib := 0
	for i := 0; i <= n; i++ {
		switch len(fibSlice) {
		case 0:
			fibSlice = append(fibSlice, 0)
		case 1:
			fibSlice = append(fibSlice, 1)
		default:
			nextFib = fibSlice[i-1] + fibSlice[i-2]
			fibSlice = append(fibSlice, nextFib)
		}
	}
	for _, v := range fibSlice {
		fmt.Fprintln(out, v)
	}
}
