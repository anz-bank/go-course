package main

import (
	"fmt"
	"io"
	"os"
)

var outfib io.Writer = os.Stdout

// Function to printh first n Fibonacci or Negafibonacci numbers
func fib(n int) {
	switch n {
	case 0:
		fmt.Fprintln(outfib, 1)
	default:
		if n > 0 {
			n1, n2 := 1, 1
			for i := 0; i < n; i++ {
				switch i {
				case 0, 1:
					fmt.Fprintln(outfib, 1)
				default:
					fmt.Fprintln(outfib, n1+n2)
					n2, n1 = n1+n2, n2
				}
			}

		} else {
			n1, n2 := 1, -1
			for i := -1; i >= n; i-- {
				switch i {
				case -1:
					fmt.Fprintln(outfib, 1)
				case -2:
					fmt.Fprintln(outfib, -1)
				default:
					fmt.Fprintln(outfib, n1-n2)
					n2, n1 = n1-n2, n2
				}

			}
		}
	}

}

func main() {
	fib(7)
	fib(-7)
}
