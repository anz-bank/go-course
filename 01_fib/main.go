package main

import (
	"fmt"
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout

func calculateFib(n int) []int {

	//need to check for edge case to avoid array out of range
	if n < 0 {
		n = 0
	}

	f0, f1, result := 1, 1, make([]int, 0, n)

	for i := 1; i <= n; i++ {
		result = append(result, f0)
		f0, f1 = f1, f0+f1
	}
	return result
}

func fib(n int) {
	//ideally this can be printed in main but lab needs to be printed inside the func
	printResult(calculateFib(n))
}

func printResult(arrayToPrint []int) {
	for i := 0; i < len(arrayToPrint); i++ {
		fmt.Fprint(outWriter, fmt.Sprintf("%d\n", arrayToPrint[i]))
	}
}

func main() {
	fib(7)
}
