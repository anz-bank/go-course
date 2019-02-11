package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {

	for num := 0; num < 10; num++ {
		fmt.Fprintln(out, printNumber(num))
	}

}
func printNumber(x int) int {
	if x < 2 {
		return x
	}
	return printNumber(x-1) + printNumber(x-2)

}
