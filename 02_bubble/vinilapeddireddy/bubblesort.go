package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// returns a sorted copy of int slice items using Bubble sort:
func bubble(items []int) []int {

	n := len(items)

	for i := 0; i < n; i++ {
		for j := n - 1; j >= i+1; j-- {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
	return items
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
