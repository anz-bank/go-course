package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//BubbleSort is a sort function for sorting an array of int
func bubbleSort(s []int) []int {
	numOfElements := len(s)
	for i := 1; i < numOfElements; i++ {
		for j := 1; j < numOfElements-i; j++ {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}

//insertionSort is a sort function for sorting an array of int
func insertionSort(s []int) []int {
	numOfElements := len(s)
	for i := 1; i < numOfElements; i++ {
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
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertionSort([]int{3, 2, 1, 5}))
}
