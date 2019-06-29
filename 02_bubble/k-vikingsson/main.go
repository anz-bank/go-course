package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func isSorted(s []int) bool {
	if len(s) <= 1 {
		return true
	}
	for index := range s[1:] {
		if s[index+1] < s[index] {
			return false
		}
	}
	return true
}

func findMin(s []int) int {
	minIndex := 0
	for index, element := range s[1:] {
		if element < s[minIndex] {
			minIndex = index + 1
		}
	}
	return minIndex
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func bubble(s []int) []int {
	copied := make([]int, len(s))
	copy(copied, s)
	for !isSorted(copied) {
		for index, element := range copied[1:] {
			if copied[index] > copied[index+1] {
				copied[index+1] = copied[index]
				copied[index] = element
			}
		}
	}
	return copied
}

func insertion(s []int) []int {
	sorted := []int{}
	copied := make([]int, len(s))
	copy(copied, s)
	for len(copied) > 0 {
		minIndex := findMin(copied)
		sorted = append(sorted, copied[minIndex])
		copied = remove(copied, minIndex)
	}
	return sorted
}

func quick(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	var lt []int
	var gt []int
	var eq []int
	pivot := s[0]
	for _, element := range s {
		switch {
		case element > pivot:
			gt = append(gt, element)
		case element < pivot:
			lt = append(lt, element)
		default:
			eq = append(eq, element)
		}
	}
	gtSorted := quick(gt)
	ltSorted := quick(lt)
	res := append(ltSorted, eq...)
	res = append(res, gtSorted...)
	return res
}

func main() {
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}
