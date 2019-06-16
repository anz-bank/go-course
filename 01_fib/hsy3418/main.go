package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//fib is a function to print out a list of fibonacci number
func fib(n int) {
<<<<<<< HEAD

	first := 1
	second := 1
	for i := 0; i < n; i++ {
		switch i {
		case 0:
			fmt.Fprintln(out, first)
		case 1:
			fmt.Fprintln(out, second)
		default:
			current := first + second
			first = second
			second = current
			fmt.Fprintln(out, current)
		}

=======
	first, second := 1, 1
	for i := 1; i <= n; i++ {
		fmt.Fprintln(out, first)
		first, second = second, first+second
>>>>>>> c86e684... rewrite fibonacci function
	}

}

func main() {
	fib(7)
}
