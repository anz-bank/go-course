package main

import (
	"fmt"
)

func fib(n int) []int {
	f := make([]int, n)
	if n < 2 {
		f = f[0:2]
	}
	f[0], f[1] = 1, 1
	for i := 2; i <= n-1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f
}

func main() {
	fmt.Println(fib(7))
}
