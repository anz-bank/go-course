package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func fib(i int) int {
	var a = 0
	var b = 1
	fmt.Println(b)
	var sum int
	for cnt := 0; cnt < i-1; cnt++ {
		sum = a + b
		fmt.Println(sum)
		a = b
		b = sum
	}
	return sum
}

func main() {
	fmt.Fprintln(out, "Hallo du schöne Welt2!")
	fib(7)
}
