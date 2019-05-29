package main

import "fmt"

const (
	fibNegaone = 1
	fibZero    = 0
	fibOne     = 1
	fibTwo     = 1
)

func fibTailNega(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	switch n {
	case 0:
		return fibZero
	case -1:
		return fibNegaone
	default:
		return fibTailNega(n+1, second, first-second )
	}
}

func fibTail(n int64, first int64, second int64) int64 {
	fmt.Println(second)
	switch n {
	case 1:
		return first
	case 2:
		return second
	default:
		return fibTail(n-1, second, first+second)
	}
}

func fib(n int) int64 {
	if n <= 0 {
		fmt.Println(fibZero)
		return fibTailNega(int64(n), fibZero, fibNegaone)
	}

	if n > 0{
		fmt.Println(fibOne)
		return fibTail(int64(n), fibOne, fibTwo)
	}

	// add this to pass golint check...
	return 0
}

func main()  {
	fib(7)
	fmt.Println("********************")
	fib(-7)
}