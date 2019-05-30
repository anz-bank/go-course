package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fibTailNega(n int64, first int64, second int64) int64 {
	fmt.Fprintln(out, second)
	switch n {
	case -1:
		return 1
	default:
		return fibTailNega(n+1, second, first-second)
	}
}

func fibTail(n int64, first int64, second int64) int64 {
	fmt.Fprintln(out, second)
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
	if n < 0 {
		return fibTailNega(int64(n), 0, 1)
	}

	if n > 0 {
		if n > 1 {
			fmt.Fprintln(out, 1)
		}

		return fibTail(int64(n), 1, 1)
	}

	// add this to pass golint check...
	fmt.Fprintln(out, 0)
	return 0
}

func main() {
	fib(7)
}
