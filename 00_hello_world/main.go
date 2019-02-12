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

func fib(n uint) {
	if n == 0 {
		return
	}
	f1,f2 := 0,1
	fmt.Fprintln(out, f2)
	var counter uint = 1
	for ; counter < n; counter++ {
		next := f1 + f2
		fmt.Fprintln(out, next)
		f1, f2 = f2, next
	}
}
