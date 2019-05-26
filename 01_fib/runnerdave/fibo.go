package fibo

import "fmt"

func fib(n int) int {
	cache := make(map[int]int)
	if n <= 1 {
		cache[0] = n
		return n
	} else if cacheVal, ok := cache[n]; ok {
		return cacheVal
	} else {
		val := fib(n-1) + fib(n-2)
		cache[n] = val
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
