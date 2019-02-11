package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, "Hallo du schöne Welt!")
	fibonacci(10)
}

func fibonacci(n int) {
	var sum int = 1
	var prev int = 0
	var cur int = 1
	for i := 1; i <= n; i++ {
		fmt.Println(sum)
		sum = prev + cur
		prev = cur
		cur = sum
	}
}
