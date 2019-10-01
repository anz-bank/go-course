package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	var num = 1
	var prev = 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Fprintln(out, num)
		tmp = prev
		prev = num
		num += tmp
	}
}

func main() {
	fib(7)
}
