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
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(sorted); i++ {
			if sorted[i] < sorted[i-1] {
				swapped = true
				sorted[i-1], sorted[i] = sorted[i], sorted[i-1]
			}
		}
	}
	return sorted
}

func insertion(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	for j := 1; j < len(sorted); j++ {
		i := j - 1
		inserted := sorted[j]
		for i >= 0 && sorted[i] > inserted {
			sorted[i+1] = sorted[i]
			i--
		}
		sorted[i+1] = inserted
	}
	return sorted
}

func main() {
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}
