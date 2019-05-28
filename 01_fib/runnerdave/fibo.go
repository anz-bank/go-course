package fibo

import "fmt"

var cache = make(map[int]int)

func fib(n int) int {
	if n == 1 || n == 0 {
		//print 1 once
		if _, ok := cache[1]; !ok {
			fmt.Println(1)
		}
		cache[n] = n
		return n
	} else if cacheVal, ok := cache[n]; ok {
		return cacheVal
	} else if n < 0 {
		sign := 1
		if n%2 == 0 {
			sign = -1
		}
		nval := sign * fib(-n)
		cache[n] = nval
		fmt.Println(nval)
		return nval
	} else {
		val := fib(n-1) + fib(n-2)
		cache[n] = val
		fmt.Println(val)
		return val
	}
}
