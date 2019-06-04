package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	sliceCopy := append(s[:0:0], s...)
	sliceLength := len(s)
	sorted := false

	for !sorted {
		swapOccurred := false
		for i := 0; i < sliceLength-1; i++ {
			if sliceCopy[i+1] < sliceCopy[i] {
				sliceCopy[i], sliceCopy[i+1] = sliceCopy[i+1], sliceCopy[i]
				swapOccurred = true
			}
		}
		if !swapOccurred {
			sorted = true
		}
	}
	return sliceCopy
}

func insertionSort(s []int) []int {
	sliceCopy := append(s[:0:0], s...)
	for key, value := range s {

		for key > 0 && sliceCopy[key-1] > value {
			sliceCopy[key] = sliceCopy[key-1]
			key--
		}
		sliceCopy[key] = value

	}
	return sliceCopy
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
