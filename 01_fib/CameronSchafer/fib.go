package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//base fibonacci function
func fib(n int) {
	fmt.Fprintln(out, n)
}

func main() {
	fib(7)
}
