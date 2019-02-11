package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fib(5)

}

func fib(n int) int {
	n1, n2 := 0, 1
	for i := 0; i < n; i++ {
		n1, n2 = n2, n1+n2
		fmt.Fprint(out, " ", n1)
	}
	return n1
}
