package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {
	const (
		a = 0
		b = 1
	)
	result := []int{a, b}

	switch {
	// take into account upper limit of what can fit into 64 bit int
	case n > 92:
		fmt.Fprintf(out, "Overflow error. Please use an argument less than 92")
	case n == 0:
		fmt.Fprintf(out, "0")
	default:
		for i := 1; i < n; i++ {
			r := len(result)
			c := result[r-2]
			d := result[r-1]
			result = append(result, c+d)
		}
		// using the for-each range loop described here: https://yourbasic.org/golang/for-loop/
		for _, v := range result[1:] {
			fmt.Fprintf(out, "%d\n", v)
		}
	}
}

func main() {
	fib(7)
}
