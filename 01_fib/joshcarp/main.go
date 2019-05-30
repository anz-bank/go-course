package main

import "fmt"

func main() {
	fib(7)
}

func fib(n int) {
	current := 1
	next := 1
	for i := 0; i < n; i++ {
		fmt.Println(current)
		current, next = next, current+next
	}
}
