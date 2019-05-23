package fibo

import "fmt"

func fib(n int) int {
	val := 0
	if n == 0 {
		return 0
	} else if n == 1 {
		val = 1
		fmt.Println(val)
		return val
	} else {
		val = fib(n-1) + fib(n-2)
		fmt.Println(val)
		return val
	}
}

// func fib(n int) int {
// 	val := 0
// 	if n == 0 {
// 		return 0
// 	} else if n == 1 {
// 		return 1
// 	} else {
// 		return fib(n-1) + fib(n-2)
// 	}
// }
