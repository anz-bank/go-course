package main

import (
	"fmt"
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout

func bubble(sliceToSort []int) []int {
	sliceLen := len(sliceToSort)

	for i := 0; i < sliceLen-1; i++ {
		for j := 0; j < sliceLen-i-1; j++ {
			if sliceToSort[j] > sliceToSort[j+1] {
				t := sliceToSort[j]
				sliceToSort[j] = sliceToSort[j+1]
				sliceToSort[j+1] = t
			}
		}
	}
	return sliceToSort
}

func insertion(sliceToSort []int) []int {
	sliceLen := len(sliceToSort)
	for i := 1; i < sliceLen; i++ {
		key := sliceToSort[i]
		j := i - 1

		for j >= 0 && sliceToSort[j] > key {
			sliceToSort[j+1] = sliceToSort[j]
			j = j - 1
		}
		sliceToSort[j+1] = key
	}
	return sliceToSort
}

func main() {
	arrayToSorted := []int{3, 2, 1, 5}
	sortedArray := bubble(arrayToSorted)
	fmt.Fprint(outWriter, sortedArray)
}
