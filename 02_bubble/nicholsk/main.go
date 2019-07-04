package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {

	// track whether the sort is complete
	sorted := false

	// make a slice copy to perform sort on
	sCopy := s

	for !sorted {

		// track whether any swaps took place in current iteration
		swaps := false

		// iterate every index in the list, starting at index 1
		for i := 1; i < len(s); i++ {

			// compare element to previous
			if sCopy[i-1] > sCopy[i] {

				// if less than previous then swap elements
				sCopy[i], sCopy[i-1] = sCopy[i-1], sCopy[i]
				swaps = true
			}

		}

		// an iteration with no swaps indicates the array is sorted
		if !swaps {
			sorted = true
		}
	}

	return sCopy
}
