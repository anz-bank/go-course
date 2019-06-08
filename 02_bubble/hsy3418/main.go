package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

/*
 Basic bubble sort for an array of int
*/
func bubbleSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := 1; j < len(s); j++ {
			if s[j] < s[j-1] {
				//swap
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}

/**
Basic insert function
**/
func insertSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		key := s[i]
		j := i - 1

		for j >= 0 && key < s[j] {
			s[j+1] = s[j]
			j = j - 1
		}
		s[j+1] = key
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5, -1, 3, 2, 3, 241}))
	fmt.Fprintln(out, insertSort([]int{3, 2, 1, 5, -1, 3, 2, 3, 241}))
}
