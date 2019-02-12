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

func fib(limit int) {
	firstNo, secondNo := 0, 1
	for i := 0; i < limit; i++ {
		fmt.Fprintln(out, firstNo)
		firstNo, secondNo = secondNo, secondNo+firstNo
	}
}
