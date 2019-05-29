package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
