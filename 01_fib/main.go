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
	var n1, n2, n3 = 0, 1, 0
	for i := 0; i < n; i++ {
		n1 = n2
		n2 = n3
		n3 = n1 + n2
		fmt.Fprint(out, n3, " ")
	}
	return n3
}
