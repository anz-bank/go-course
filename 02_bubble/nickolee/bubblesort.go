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

	unsorted := true
	for unsorted {
		// perform the check to see if sorting is complete to exit for loop
		unsorted = false
		for i := len(result) - 1; i > 0; i-- {
			if result[i] < result[i-1] {
				result[i], result[i-1] = result[i-1], result[i]
				unsorted = true
			}
		}
	}
	return result
}

func main() {
	s := []int{3, 2, 1, 5}
	fmt.Fprintln(out, bubbleSort(s))
}
