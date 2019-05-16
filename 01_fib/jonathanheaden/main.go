package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) string {
	f := [3]int{0, 1, 1}
	var prefix, res string
	var j int
	var neg = false
	if n < 0 {
		n *= -1
		neg = true
	}
	for i := 0; i < n; i++ {
		if j = i + 1; i > 1 {
			f = [3]int{f[1], f[2], (f[1] + f[2])}
			j = 2
		}
		if prefix = ""; (i%2 == 1) && neg {
			prefix = "-"
		}
		res = fmt.Sprintf("%s\n%s%d", res, prefix, f[j])
	}
	return res
}

func main() {

	fmt.Fprintln(out, fib(7))

}
