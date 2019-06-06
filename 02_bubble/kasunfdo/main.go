package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	out := make([]int, len(s))
	copy(out, s)

	for i := len(out); i > 0; i-- {
		swapped := false
		for j := 1; j < i; j++ {
			if out[j-1] > out[j] {
				out[j-1], out[j] = out[j], out[j-1]
				swapped = true
			}
		}
		// End looping if current array is sorted
		if !swapped {
			break
		}
	}

	return out
}

func insertionSort(input []int) []int {
	out := make([]int, len(input))
	copy(out, input)

	for i := 1; i < len(out); i++ {
		for j := i; j > 0; j-- {
			if out[j-1] > out[j] {
				out[j-1], out[j] = out[j], out[j-1]
			}
		}
	}

	return out
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
