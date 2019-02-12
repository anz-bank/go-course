package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(items []int) []int {
	n := len(items)
	sorted := false

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n -= 1
	}
	return items
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
