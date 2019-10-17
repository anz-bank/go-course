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
	n1, n2 := 1, 1
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, n1)
		n1, n2 = n2, n1+n2
	}
}
