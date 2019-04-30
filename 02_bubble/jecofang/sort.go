package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	n := len(s)
	arr := make([]int, n)
	copy(arr, s)

	if n < 2 {
		return arr
	}

	swap := true
	for swap {
		swap = false
		for i := 0; i < n-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swap = true
			}
		}
	}

	return arr
}

func insertion(s []int) []int {
	n := len(s)
	arr := make([]int, n)
	copy(arr, s)

	if n < 2 {
		return arr
	}

	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	return arr
}

func quick(s []int) []int {
	n := len(s)

	if n < 2 {
		arr := make([]int, n)
		copy(arr, s)
		return arr
	}

	left := make([]int, 0, n)
	middle := make([]int, 0, n)
	right := make([]int, 0, n)
	pivot := s[n-1]

	for _, v := range s {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		case v > pivot:
			right = append(right, v)
		}
	}

	left, right = quick(left), quick(right)

	left = append(left, middle...)
	left = append(left, right...)

	return left
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
