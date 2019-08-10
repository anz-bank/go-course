package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// Bubble implements the Bubble sort algorithm.
func bubble(s []int) []int {
	length := len(s)
	// Create copy
	result := make([]int, length)
	copy(result, s)
	// Increment through the slice
	// All but last number need to be checked
	for i := 1; i < length; i++ {
		// Increment through the slice
		// Compare current index with next
		// Don't check previous sorted indices
		for j := 0; j < length-i; j++ {
			// Swap unsorted values
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

// Insertion implements the Insertion sort algorithm.
func insertion(s []int) []int {
	length := len(s)
	// Create copy
	result := make([]int, length)
	copy(result, s)
	// Increment through the slice
	// All but first number need to be checked
	for i := 1; i < length; i++ {
		// Increment backwards through slice
		// Start with first unsorted index
		for j := i; j > 0; j-- {
			// Swap unsorted values
			if result[j] < result[j-1] {
				result[j], result[j-1] = result[j-1], result[j]
			}
		}
	}

	return result
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprint(out, insertion([]int{3, 2, 1, 5}))
}
