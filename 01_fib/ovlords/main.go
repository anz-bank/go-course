package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(n int) {

	lines := 1
	currentNumber := 1
	previousNumber := 0
	scratchNumber := 0

	for lines <= n {

		fmt.Fprintln(out, currentNumber)

		scratchNumber = previousNumber
		previousNumber = currentNumber
		currentNumber += scratchNumber

		lines++
	}

}

func main() {
	fib(7)
}
