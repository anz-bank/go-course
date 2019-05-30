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
				temp := s[j]
				s[j] = s[j-1]
				s[j-1] = temp
			}
		}
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
}
