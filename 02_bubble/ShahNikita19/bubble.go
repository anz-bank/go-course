package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
}

func bubbleSort(array []int) []int {
	arrayLength := len(array)
	copyArray := make([]int, len(array))
	copy(copyArray, array)
	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
