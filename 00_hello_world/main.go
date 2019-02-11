package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	fibonacciSeries(7)
}

func fibonacciSeries(number int) {
	var firstNumber = 1
	var secondNumber = 1
	fmt.Fprintln(out, firstNumber)
	fmt.Fprintln(out, secondNumber)
	for count := 3; count <= number; count++ {
		var sum = firstNumber + secondNumber
		fmt.Fprintln(out, sum)
		firstNumber = secondNumber
		secondNumber = sum

	}
}
