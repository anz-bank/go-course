package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

//bubbleSort compares two adjacent numbers
func bubbleSort(s []int) []int {
	for i := len(s); i > 0; i-- {
		for j := 1; j < i; j++ {
			if s[j-1] >= s[j] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}

func insertSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		var j int
		for j = i; j > 0; j-- {
			if s[j-1] <= s[i] {
				break // We've found the insertion point
			}
		}
		// insert s[i] at s[j] moving the rest up
		tmp := s[i]
		copy(s[j+1:i+1], s[j:i])
		s[j] = tmp
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertSort([]int{3, 2, 1, 5}))
}
