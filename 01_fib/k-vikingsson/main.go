package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func negaFib(n int) {
	c := 0
	n0 := 0
	n1 := 1
	nC := n0 - n1
	for c >= n {
		switch c {
		case 0:
			fmt.Fprintln(out, 0)
		case -1:
			fmt.Fprintln(out, n1)
		default:
			fmt.Fprintln(out, nC)
			n0 = n1
			n1 = nC
			nC = n0 - n1
		}
		c--
	}
}

func fib(n int) {
	if n <= 0 {
		negaFib(n)
		return
	}
	c := 1
	nA := 1
	nB := 1
	nC := nA + nB
	for c <= n {
		switch c {
		case 1:
			fmt.Fprintln(out, nA)
		case 2:
			fmt.Fprintln(out, nB)
		default:
			fmt.Fprintln(out, nC)
			nA = nB
			nB = nC
			nC = nA + nB
		}
		c++
	}
}

func main() {
	fib(7)
}
