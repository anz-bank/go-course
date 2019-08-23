package main

import (
	"fmt"
)

// fibonacci returns a function that returns
// an number in the fibonacci sequence.
// This is an implmentation of "closure".
// The inside function returns the current number
// in the fibonnacci sequence, starting at 1
// and sets up the next which will be returned
// in the subsequent call
func fibonacci() func() int {
	secondLast, last := 0, 1

	return func() int {
		result := last
		secondLast, last = last, secondLast+last
		return result
	}
}
func fib(n int) {
	f := fibonacci()

	for i := 0; i < n; i++ {
		fmt.Println(f())
	}

}
func main() {
	fib(7)
}
