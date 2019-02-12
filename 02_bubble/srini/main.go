package main

import (
	"fmt"
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout

func bubble(sliceToSort []int) []int {
	sortedSlice := make([]int, len(sliceToSort))
	copy(sortedSlice, sliceToSort)

	lastIndex := len(sortedSlice) - 1

	for i := 0; i < lastIndex; i++ {
		swapped := true
		for j := 0; j < lastIndex-i; j++ {
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j+1], sortedSlice[j] = sortedSlice[j], sortedSlice[j+1]
				swapped = false
			}
		}
		// we can stop once there is no sorting in inner loop
		if swapped {
			break
		}
	}
	return sortedSlice
}

func insertion(sliceToSort []int) []int {
	sortedSlice := make([]int, len(sliceToSort))
	copy(sortedSlice, sliceToSort)

	sliceLen := len(sortedSlice)
	for i := 1; i < sliceLen; i++ {
		key := sortedSlice[i]
		j := i - 1

		for ; j >= 0 && sortedSlice[j] > key; j-- {
			sortedSlice[j+1] = sortedSlice[j]
		}
		sortedSlice[j+1] = key
	}
	return sortedSlice
}

func main() {
	arrayToSort := []int{3, 2, 1, 5}
	sortedArray := bubble(arrayToSort)
	fmt.Fprint(outWriter, sortedArray)
}
