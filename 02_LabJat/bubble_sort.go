package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubbleSort(input []int) []int {
	arLength := len(input)
	for i := 0; i < arLength; i++ {

		for j := 0; j < arLength-i-1; j++ {
			if input[j] > input[j+1] {
				input[j+1], input[j] = input[j], input[j+1]
			}
		}

	}
	return input
}
func insertionSort(input []int) []int {
	arLength := len(input)
	var j int
	for i := 0; i < arLength; i++ {
		key := input[i]

		for j = i - 1; (j >= 0) && input[j] > key; j-- {
			input[j+1] = input[j]
		}
		input[j+1] = key
	}
	return input
}

func main() {
	fmt.Fprintln(out, bubbleSort([]int{3, 2, 1, 5}))
}
