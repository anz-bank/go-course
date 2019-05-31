package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {

	// Copy slice: See https://github.com/go101/go101/wiki
	sorted := append(s[:0:0], s...)

	for range sorted {
		noSwaps := true

		// need to do n-1 compare/swaps
		for i := range sorted[1:] {

			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}

	}
	return sorted
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
