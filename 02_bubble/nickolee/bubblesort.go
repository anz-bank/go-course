package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// general sorting algorithm that sorts by "bubbling" smaller numbers to the left
func bubbleSort(s []int) []int {
	result := make([]int, len(s))
	copy(result, s)
	sorted := false

	// perform the check to see if sorting is complete to exit for loop
	for !sorted {
		sorted = true
		for i := len(result) - 1; i > 0; i-- {
			if result[i] < result[i-1] {
				result[i], result[i-1] = result[i-1], result[i]
				sorted = false
			}
		}
	}
	return result
}

func main() {
	s := []int{3, 2, 1, 5}
	fmt.Fprintln(out, bubbleSort(s))
}
