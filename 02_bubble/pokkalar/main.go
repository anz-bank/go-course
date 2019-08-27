package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubbleSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	for i := 0; i < len(c); i++ {
		swapped := false
		for j := 0; j < len(c)-1; j++ {
			if c[j] > c[j+1] {
				c[j+1], c[j] = c[j], c[j+1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return c
}

func insertionSort(s []int) []int {
	i := make([]int, len(s))
	copy(i, s)
	for j := 1; j < len(i); j++ {
		for k := j; k > 0; k-- {
			if i[k] < i[k-1] {
				i[k-1], i[k] = i[k], i[k-1]
			}
		}
	}
	return i
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
}
