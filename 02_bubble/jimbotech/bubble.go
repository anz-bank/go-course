package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)

	for k := range sorted {
		for i := range sorted[:len(sorted)-k-1] {
			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
			}
		}
	}
	return sorted
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
