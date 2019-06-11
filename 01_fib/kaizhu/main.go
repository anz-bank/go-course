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

func fib(n int64) {
	switch {
	case n == 0:
		fmt.Fprintln(out, 0)
	case n < -92 || n > 92:
		fmt.Fprintln(out, "Please enter a number within the range [-92, 92].")
	default:
		var sign int64 = 1
		if n < 0 {
			n = -n
			sign = -1
		}
		var a, b int64 = 1, sign
		for i := int64(0); i < n; i++ {
			fmt.Fprintln(out, a)
			a, b = b, a+b*sign
		}
	}
}
