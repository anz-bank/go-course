package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	printfib(7)
}

func printfib(n int) {
	// as per specs: intentionally skip 0
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, fibN(i))
	}
}

func fibN(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return (fibN(n-1) + fibN(n-2))
	}
}
