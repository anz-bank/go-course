package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, strings.Trim(strings.Join(strings.Split(fmt.Sprint(fib(7)), " "), "\n"), "[]"))
}

func fib(n int) []int {
	if n == 0 {
		return []int{0}
	}

	isNegativeFib := false

	if n < 0 {
		isNegativeFib = true
		n *= -1
	}

	result := []int{1}
	for i := 1; i < n; i++ {
		var previousValue int
		if len(result) == 1 {
			previousValue = 0
		} else {
			previousValue = result[len(result)-2]
		}

		if isNegativeFib {
			result = append(result, previousValue-result[i-1])
		} else {
			result = append(result, result[i-1]+previousValue)
		}
	}
	return result
}
