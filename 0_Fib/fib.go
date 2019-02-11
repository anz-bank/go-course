package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func Fibonacci(n int) []int {

	fib, t1, t2 := []int{}, 1, 1
	for i := 0; i < n; i++ {
		// add the next member
		fib = append(fib, t1)
		t1, t2 = t2, t1+t2
	}
	return fib
}
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func main() {
	fmt.Fprint(out, "Fibonnaci Series of 7: ", arrayToString(Fibonacci(7), "-"))

}
