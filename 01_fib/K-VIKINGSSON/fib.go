package main

import (
	"fmt";
)

func negaFib(n int) {
	c := 0
	n0 := 0
	n1 := 1
	nC := n0 - n1
	for c >= n {
		if c == 0 {
			fmt.Println(0)
		} else if c == -1 {
			fmt.Println(n1)
		} else {
			fmt.Println(nC)
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
		if c == 1 {
			fmt.Println(nA)
		} else if c == 2 {
			fmt.Println(nB)
		} else {
			fmt.Println(nC)
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
