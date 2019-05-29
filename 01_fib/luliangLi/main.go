package main

import "fmt"

const (
	fibNegaone = 1
	fibZero    = 0
	fibOne     = 1
	fibTwo     = 1
)

func fib_tail_nega(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	switch n {
	case 0:
		return fibZero
	case -1 :
		return fibNegaone
	default :
		return fib_tail_nega(n + 1, second, first - second )
	}
}

func fib_tail(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	switch n {
	case 1 :
		return first
	case 2 :
		return second
	default :
		return fib_tail(n-1, second, first + second)
	}
}

func fib(n int) int64 {
	if n <= 0 {
		fmt.Println(fibZero)
		return fib_tail_nega(int64(n), fibZero, fibNegaone)
	}

	if n > 0{
		fmt.Println(fibOne)
		return fib_tail(int64(n), fibOne, fibTwo)
	}

	// add this to pass golint check...
	return 0
}

func main()  {
	fib(7)
	fmt.Println("********************")
	fib(-7)
}