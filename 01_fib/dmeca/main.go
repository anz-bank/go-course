package main

import "fmt"

func main() {
	fib(7)
	fmt.Println("------")
	fib(-7)
}

func fib(n int) {
	var a, b, c int = 0, 1, 0
	var posneg int = 1

	if n < 0 {
		posneg = -1
		n *= -1
	}

	for idx := 1; idx <= n; idx++ {
		fmt.Println(b)
		c = a + b*posneg
		a = b
		b = c
	}
}
