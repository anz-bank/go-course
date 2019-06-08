package main

import (
	"fmt"
)

func main() {
	fib(7)
}

func fib(n int) {
	current := 0
	next := 1
	sign := 1
	if n < 0 {
		sign = -1
		n *= -1
	}
	nextSign := sign
	for i := 0; i < n; i++ {
		sign *= nextSign
		current, next = next, current+next
		fmt.Println(current * sign)
	}
}
