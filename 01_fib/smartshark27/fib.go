package main

import (
	"fmt"
)

func fib(n int) {
	var num = 1
	var prev = 0
	var tmp int
	for i := 0; i < n; i++ {
		fmt.Println(num)
		tmp = prev
		prev = num
		num += tmp
	}
}

func main() {
	fib(7)
}