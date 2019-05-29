package main

import "fmt"

const (
	fib_NEGAONE = 1
	fib_ZERO = 0
	fib_ONE = 1
	fib_TWO = 1
)

func fib_tail_nega(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	if n == 0 {
		return fib_ZERO
	} else if n == -1 {
		return fib_NEGAONE
	} else {
		return fib_tail_nega(n + 1, second, first - second )
	}
}

func fib_tail(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	if n == 1 {
		return first
	} else if n == 2 {
		return second
	} else {
		return fib_tail(n-1, second, first + second)
	}
}

func fib(n int) int64 {
	if n <= 0 {
		fmt.Println(fib_ZERO)
		return fib_tail_nega(int64(n), fib_ZERO, fib_NEGAONE)
	} else {
		fmt.Println(fib_ONE)
		return fib_tail(int64(n), fib_ONE, fib_TWO)
	}
}

func main()  {
	fib(7)
	fib(-7)
}