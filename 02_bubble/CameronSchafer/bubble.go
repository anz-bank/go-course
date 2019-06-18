package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	// requirement was to sort a copy of the array.
	t := make([]int, len(s))
	copy(t, s)
	check := false

	for !check {
		swapped := true
		// sort once and then check
		for i := 0; i < len(t); i++ {
			// check if in last slice
			if i+1 < len(t) {
				if t[i] > t[i+1] {
					// swap the elems around
					t[i], t[i+1] = t[i+1], t[i]
					swapped = false
				}
			}
		}
		check = swapped
	}
	return t
}

func insertion(s []int) []int {
	// loop through the given slice
	for i := 0; i < len(s); i++ {
		a, firstPart, secondPart := s[i], s[:i], s[i+1:]

		for j := len(firstPart) - 1; j >= 0; j-- {
			if a < firstPart[j] {
				newA, b := a, firstPart[j]
				// create the new slice and reset the parent loop
				s, i = append(append(firstPart[:j], newA), append([]int{b}, secondPart...)...), 0
				break
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}
