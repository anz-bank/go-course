package main

import (
	"fmt"
)

func main() {
	fib(-7)

}

func fib(n int) {
	current := 1
	next := 1
	sign := 1
	if n < 0 {
		sign = -1
		n *= -1
	}
	nextSign := sign
	for i := 0; i < n; i++ {
		sign *= nextSign
		fmt.Println(current * sign)
		current, next = next, current+next
	}
}
