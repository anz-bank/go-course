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

func fib(num int) {
	num1 := 0
	num2 := num1 + 1
	var result string
	for i := 0; i < num; i++ {
		result += strconv.Itoa(num1)
		temp := num2
		num2 += num1
		num1 = temp
	}
	fmt.Fprintln(out, result)
}
