package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var outWriter io.Writer = os.Stdout

func calculateFib(n int) []int {
	f0, f1, result := 1, 1, []int{}

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
	resultsAsStr := []string{}
	for i := 0; i < len(arrayToPrint); i++ {
		resultsAsStr = append(resultsAsStr, strconv.Itoa(arrayToPrint[i]))
	}
	fmt.Fprintln(outWriter, strings.Join(resultsAsStr, "\n"))
}

func main() {
	fib(7)
}
