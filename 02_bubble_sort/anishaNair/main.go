package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(n []int) []int {
	for i := 0; i < len(n); i++ {
		for j := 0; j < ((len(n) - i) - 1); j++ {
			if n[j] > n[j+1] {
				n[j+1], n[j] = n[j], n[j+1]
			}
		}
	}
	return n
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
