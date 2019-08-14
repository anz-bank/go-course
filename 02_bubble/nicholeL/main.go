package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//Compare two adjacent numbers
func bubbleSort(s []int) []int {
	for index := 0; index < len(s); index++ {
		for count := 0; count < len(s)-index-1; count++ {
			if s[count] >= s[count+1] {
				s[count], s[count+1] = s[count+1], s[count]
			}
		}
	}
	return s
}

func insertSort(s []int) []int {

	for first := 1; first < len(s); first++ {
		index := first
		temp := s[first]
		for index > 0 && temp < s[index-1] {
			s[index] = s[index-1]
			index--
		}
		s[index] = temp
	}

	return s
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertSort([]int{3, 2, 1, 5}))
}
