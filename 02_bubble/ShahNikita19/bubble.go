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
	copyArray := make([]int, len(array))
	copy(copyArray, array)
	for i := 0; i < len(copyArray); i++ {
		for j := 0; j < len(copyArray)-i-1; j++ {
			if copyArray[j] > copyArray[j+1] {
				copyArray[j], copyArray[j+1] = copyArray[j+1], copyArray[j]
			}
		}
	}
	return copyArray
}
