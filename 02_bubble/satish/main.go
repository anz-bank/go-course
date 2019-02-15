package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	// create copy of the input
	cp := make([]int, len(s))
	copy(cp, s)

	// length of the array
	n := len(cp)

	for swapped := true; swapped; {
		swapped = false

		for i := 1; i < n; i++ {
			if cp[i-1] > cp[i] {
				fmt.Println("Swapping in progress...")
				// leveraging Go's tuple assignment
				cp[i], cp[i-1] = cp[i-1], cp[i]
				// set swapped to true - if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is fully sorted.
				swapped = true
			}
		}
	}
	return cp
}

func insertion(s []int) []int {
	// create copy of the input
	cp := make([]int, len(s))
	copy(cp, s)

	// length of the array
	n := len(s)

	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if cp[j-1] > cp[j] {
				cp[j-1], cp[j] = cp[j], cp[j-1]
			}
			j--
		}
	}
	return cp
}

func main() {
	a := []int{3, 2, 1, 5}
	fmt.Fprintln(out, bubble(a))
}
