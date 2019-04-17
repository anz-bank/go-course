package fib

import (
	"fmt"
)

func Fib(n int) {
	if n == 0 {
		return
	} else if n == 1 {
		fmt.Println(1)
		return
	}

	fibResult := make([]int, n)

	fibResult[0] = 1
	fibResult[1] = 1

	for i := 2; i < n; i++ {
		fibResult[i] = fibResult[i-1] + fibResult[i-2]
	}

	for j := 0; j < n; j++ {
		fmt.Println(fibResult[j])
	}
}
