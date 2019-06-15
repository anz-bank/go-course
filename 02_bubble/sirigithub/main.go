package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubbleSort(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}

	return s
}

func insertionSort(s []int) []int {
	i := 1

	for i < len(s) {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
		i++
	}
	return s

}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))

}
