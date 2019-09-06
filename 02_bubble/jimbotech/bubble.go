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

	var ss []int
	for k := range sorted {
		ss = sorted[:len(sorted)-k-1]
		for i, val := range ss {
			if val > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
			}
		}
	}
	return sorted
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
