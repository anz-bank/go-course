package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, fibonacci(10))
}

func fibonacci(n int) string {
	var fibSeries string
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		fibSeries += strconv.Itoa(a)
		a, b = b, a+b
	}
	return fibSeries
}
