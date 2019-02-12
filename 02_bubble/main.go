package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(s)-1; i++ {
			if s[i+1] < s[i] {
				swapped = true
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
	}
	return s
}

func insertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j] > s[j+1] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
	return s
}
