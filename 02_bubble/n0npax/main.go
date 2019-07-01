package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

// partition divides slice into partitions based on pivot high
func partition(s []int, low int, high int) int {
	pivot, i := s[high]+1, low-1
	for j := low; j < high; j++ {
		if s[j] < pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}
	s[i+1], s[high] = s[high], s[i+1]
	return i + 1
}

// qsort sorts provided partition. Modifies reference
func qsort(s []int, lo int, hi int) {
	if lo < hi {
		// p - partition vertical pivot
		p := partition(s, lo, hi)
		qsort(s, lo, p-1)
		qsort(s, p+1, hi)
	}

}

// quicksort returns sorted list(copy) using quicksort
// quicksort has average compexity O(nlogn)
func quicksort(s []int) []int {
	ts := append(s[:0:0], s...)
	qsort(ts, 0, len(ts)-1)
	return ts
}

// insertion sorts list(copy) using insertion
func insertion(s []int) []int {
	ts := append(s[:0:0], s...)

	for i := range ts {
		for i > 0 {
			if ts[i-1] > ts[i] {
				ts[i-1], ts[i] = ts[i], ts[i-1]
			}
			i--
		}
	}
	return ts
}

// bubble sorts list(copy) using buble sort
func bubble(s []int) []int {
	ts := append(s[:0:0], s...)

	for i := len(ts); i > 0; i-- {
		swapped := false
		for j := 1; j < i; j++ {
			if ts[j-1] > ts[j] {
				ts[j-1], ts[j] = ts[j], ts[j-1]
				swapped = true
			}
		}
		if !swapped {
			// If no elements were swapped, the slice is sorted so we can stop here
			return ts
		}
	}

	return ts
}

// main - entrypoint provided due to lab requirements
func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
