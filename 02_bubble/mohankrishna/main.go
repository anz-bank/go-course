package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)

	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(sorted)-1; i++ {
			if sorted[i+1] < sorted[i] {
				swapped = true
				sorted[i+1], sorted[i] = sorted[i], sorted[i+1]
			}
		}
	}
	return sorted
}

func insertionSort(s []int) []int {
	sorted := make([]int, len(s))
	copy(sorted, s)
	for i := 1; i < len(sorted); i++ {
		for j := i - 1; j >= 0; j-- {
			if sorted[j] > sorted[j+1] {
				sorted[j+1], sorted[j] = sorted[j], sorted[j+1]
			}
		}
	}
	return sorted
}
