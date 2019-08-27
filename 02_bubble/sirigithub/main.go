package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubbleSort(s []int) []int {
	n := len(s)
	arrayToSort := append(s[:0:0], s...)
	for {
		newn := 0
		for i := 1; i < n; i++ {
			if arrayToSort[i-1] > arrayToSort[i] {
				arrayToSort[i-1], arrayToSort[i] = arrayToSort[i], arrayToSort[i-1]
				newn = i
			}
		}
		n = newn
		if n <= 1 {
			break
		}
	}
	return arrayToSort
}

func insertionSort(s []int) []int {
	arrayToSort := append(s[:0:0], s...)

	for i := 1; i < len(arrayToSort); i++ {
		key := arrayToSort[i]
		j := i - 1

		for j >= 0 && arrayToSort[j] > key {
			arrayToSort[j+1] = arrayToSort[j]
			j--
		}
		arrayToSort[j+1] = key
	}
	return arrayToSort
}
func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
}
