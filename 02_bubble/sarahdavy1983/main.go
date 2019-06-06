package main

import (
	"fmt"
)

func bubbleSort(s []int) []int {
	sliceCopy := append(s[:0:0], s...)
	n := len(s)
	swapped := true
	for swapped {
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
	sliceCopy := append(s[:0:0], s...)
	n := len(s)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if sliceCopy[j-1] > sliceCopy[j] {
				sliceCopy[j-1], sliceCopy[j] = sliceCopy[j], sliceCopy[j-1]
			}
			j--
		}
	}
	return sliceCopy
}

func main() {
	fmt.Println(bubbleSort([]int{3, 2, 1, 5}))
	fmt.Println(insertionSort([]int{3, 2, 1, 5}))
}
