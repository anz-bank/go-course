package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	sortedSlice := make([]int, len(s))
	copy(sortedSlice, s)
	for i := 0; i < len(sortedSlice)-1; i++ {
		for j := 0; j < len(sortedSlice)-i-1; j++ {
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j+1], sortedSlice[j] = sortedSlice[j], sortedSlice[j+1]
			}
		}
	}
	return sortedSlice
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
