package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	// length of the array
	n := len(s)

	for swapped := true; swapped; {
		swapped = false

		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				fmt.Println("Swapping in progress...")
				// leveraging Go's tuple assignment
				s[i], s[i-1] = s[i-1], s[i]
				// set swapped to true - if the loop ends and swapped is still equal
				// to false, our algorithm will assume the list is fully sorted.
				swapped = true
			}
		}
	}
	return s
}

func insertion(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
			j--
		}
	}
	return s
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprint(out, insertion([]int{3, 2, 1, 5}))
}