package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func calculateFirstNFibNumArray(n int) []uint64 {
	if n <= 0 {
		return nil
	}

	fibResults := make([]uint64, n)
	fibResults[0] = 1
	if n == 1 {
		return fibResults[:1]
	}

	fibResults[1] = 1
	for i := 2; i < n; i++ {
		fibResults[i] = fibResults[i-1] + fibResults[i-2]
	}
	return fibResults
}

func printIntSlice(intSlice []uint64) {
	for _, v := range intSlice {
		fmt.Fprintln(out, v)
	}
}

func fib(n int) {
	printIntSlice(calculateFirstNFibNumArray(n))
}

func main() {
	fib(7)
}
