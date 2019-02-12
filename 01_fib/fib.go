package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func fib(n int) string {
	f := make([]int, n)
	if n < 2 {
		f = f[0:2]
	}
	f[0], f[1] = 1, 1
	for i := 2; i <= n-1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return strings.Trim(strings.Replace(fmt.Sprint(f), " ", " ", -1), "[]")
}

func main() {
	fmt.Fprintf(out, "%v", fib(7))
}
