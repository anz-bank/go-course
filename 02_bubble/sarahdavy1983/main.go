package main

import (
	"fmt"
)

func bubbleSort(s []int) []int {
	n := len(s)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}
	}
	return s
}
func insertionSort(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
			j--
		}
	}
	return s
}

func main() {
	fmt.Println(bubbleSort([]int{3, 2, 1, 5}))
	fmt.Println(insertionSort([]int{3, 2, 1, 5, -1}))
}
