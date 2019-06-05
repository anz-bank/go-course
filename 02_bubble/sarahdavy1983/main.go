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

func main() {
	fmt.Println(bubbleSort([]int{3, 2, 1, 5}))
}
