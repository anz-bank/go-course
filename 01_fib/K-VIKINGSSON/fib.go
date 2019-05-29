package main

import (
	"fmt";
)

func negaFib(n int) {
	c := 0
	n_0 := 0
	n_1 := 1
	n_c := n_0 - n_1
	for c >= n {
		if c == 0 {
			fmt.Println(0)
		} else if c == -1 {
			fmt.Println(n_1)
		} else {
			fmt.Println(n_c)
			n_0 = n_1
			n_1 = n_c
			n_c = n_0 - n_1
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
	n_a := 1
	n_b := 1
	n_c := n_a + n_b
	for c <= n {
		if c == 1 {
			fmt.Println(n_a)
		} else if c == 2 {
			fmt.Println(n_b)
		} else {
			fmt.Println(n_c)
			n_a = n_b
			n_b = n_c
			n_c = n_a + n_b
		}
		c++
	}
}

func main() {
	fib(7)
	fib(-7)
}
