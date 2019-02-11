package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

func main() {
	fib(7)
}

func fib(limit int) {
	firstNo, secondNo := 0, 1
	var finalString string
	for i := 0; i < limit; i++ {
		finalString += strconv.Itoa(firstNo)
		firstNo, secondNo = secondNo, secondNo+firstNo
	}
	fmt.Fprintln(out, finalString)
}
