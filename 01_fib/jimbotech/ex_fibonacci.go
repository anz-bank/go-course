package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

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

func fibSeries(n int) []int {
	counter := n
	if n < 0 {
		counter = -n
	}
	var fibSerial []int
	f := fibonacci()
	for i := 0; i < counter; i++ {
		factor := 1
		if n < 0 && i%2 != 0 {
			factor = -1
		}
		fibSerial = append(fibSerial, f()*factor)
	}
	return fibSerial
}

func fib(n int) {
	fibSer := fibSeries(n)
	for _, v := range fibSer {
		fmt.Fprintln(out, v)
	}
}

func main() {
	fib(7)
}
