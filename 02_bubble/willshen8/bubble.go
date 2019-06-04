package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	sliceLength := len(s)
	sorted := false

	for !sorted {
		swapOccurred := false
		for i := 0; i < sliceLength-1; i++ {
			if s[i+1] < s[i] {
				swap(s, i, i+1)
				swapOccurred = true
			}
		}
		if !swapOccurred {
			sorted = true
		}
	}
	return s
}

func insertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		currValue := s[i]
		emptySlot := i

		for emptySlot > 0 && s[emptySlot-1] > currValue {
			s[emptySlot] = s[emptySlot-1]
			emptySlot--
		}
		s[emptySlot] = currValue

	}
	return s
}

func swap(slice []int, pos1 int, pos2 int) []int {
	temp := slice[pos2]
	slice[pos2] = slice[pos1]
	slice[pos1] = temp
	return slice
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}
