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
	swapped := true
	for swapped {
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

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
