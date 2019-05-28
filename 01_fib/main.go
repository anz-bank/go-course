package main

import "errors"

func fib(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Fibonacci number is negative")
	}
	switch n {
	case 1:
		return 1, nil
	case 0:
		return 0, nil
	default:
		fib1, _ := fib(n - 1)
		fib2, _ := fib(n - 2)
		return fib1 + fib2, nil
	}
}

func main() {
}
