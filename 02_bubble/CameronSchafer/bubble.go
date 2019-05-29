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
		tempCheck := true
		// sort once and then check
		for i := 0; i < len(t); i++ {
			// check if in last slice
			if i+1 < len(t) {
				if t[i] > t[i+1] {
					// swap the elems around
					t[i], t[i+1] = t[i+1], t[i]
					tempCheck = false
				}
			}
		}
		check = tempCheck
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

// accidentally created an O(n2) complexity instead of O(n log n) complexity
// merge sort: O(n2) sorting algorithm
func merge(s []int) []int {
	sBox := [][]int{}
	for a := range s {
		sBox = append(sBox, []int{s[a]})
	}

	for len(sBox) > 1 {
		var tBox [][]int
		for i := 0; i < len(sBox); i += 2 {
			// if the adjacent box exists then compare them
			if i+1 != len(sBox) {
				tSlice := []int{}
				// need to do for loop here to check if i and i+1 lengths are greater than 0
				for len(sBox[i]) > 0 && len(sBox[i+1]) > 0 {
					a, b := sBox[i][0], sBox[i+1][0]
					if a < b {
						tSlice = append(tSlice, a)
						sBox[i] = sBox[i][1:]
					} else {
						tSlice = append(tSlice, b)
						sBox[i+1] = sBox[i+1][1:]
					}
				}
				// need to do a check here for which array has less values
				if len(sBox[i]) > 0 {
					tSlice = append(tSlice, sBox[i]...)
				} else if len(sBox[i+1]) > 0 {
					tSlice = append(tSlice, sBox[i+1]...)
				}
				tBox = append(tBox, tSlice)
			} else {
				tBox = append(tBox, sBox[i])
			}
		}
		// set sBox to value of tBox
		sBox = tBox
	}
	return sBox[0]
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, merge([]int{3, 2, 1, 5}))
}
