package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {

	if n > 0 {

		currentNumber, previousNumber, scratchNumber := 1, 0, 0

		for lines := 1; lines <= n; lines++ {
			fmt.Fprintln(out, currentNumber)

			scratchNumber, previousNumber = previousNumber, currentNumber
			currentNumber += scratchNumber
		}

	} else {
		fmt.Fprintln(out, "Invalid Input: Must be a positive integer")
	}

}

func main() {
	fib(7)
}
