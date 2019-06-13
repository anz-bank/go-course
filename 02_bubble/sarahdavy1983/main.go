package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubbleSort(s []int) []int {

	var sliceCopy = append(s[:0:0], s...)
	n := len(s)
	for swapped := true; swapped; {
		swapped = false
		for i := 1; i < n; i++ {
			if sliceCopy[i-1] > sliceCopy[i] {
				sliceCopy[i], sliceCopy[i-1] = sliceCopy[i-1], sliceCopy[i]
				swapped = true
			}
		}
	}
	return sliceCopy
}
func insertionSort(s []int) []int {
	var sliceCopy = append(s[:0:0], s...)
	n := len(s)
	for i := 1; i < n; i++ {
		for j := i; j > 0; {
			if sliceCopy[j-1] > sliceCopy[j] {
				sliceCopy[j-1], sliceCopy[j] = sliceCopy[j], sliceCopy[j-1]
			}
			j--
		}
	}
	return sliceCopy
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertionSort([]int{3, 2, 1, 5}))
}
