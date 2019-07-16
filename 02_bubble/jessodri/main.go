package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	b := append([]int{}, s...)
	length := len(b)
	for i := length - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
	return b
}

func insertion(s []int) []int {
	b := append([]int{}, s...)
	length := len(b)
	for i := 0; i < length; i++ {
		current := b[i]
		j := 0
		for j = i - 1; j >= 0 && b[j] > current; j-- {
			b[j+1] = b[j]
		}
		b[j+1] = current
	}
	return b
}

func quicksort(s []int) []int {
	length := len(s)
	if length <= 1 {
		return s
	}
	first := s[0]
	left := []int{}
	right := []int{}

	for i := 1; i < length; i++ {
		if s[i] < first {
			left = append(left, s[i])
		} else {
			right = append(right, s[i])
		}
	}
	result := append(quicksort(left), first)
	result = append(result, quicksort(right)...)
	return result
}

func main() {
	fmt.Fprintln(out, bubble([]int{1, 3, 2, 67, 43, 4, 5}))
}
