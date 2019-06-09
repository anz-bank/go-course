package main

/*
	https://github.com/undrewb/go-course/tree/lab1/02_bubble

	Lab 2 - Bubble sort
	Create an executable go program in directory 02_bubble/USERNAME
	Write a function that returns a sorted copy of int slice s using bubble sort:
	func bubble(s []int) []int
	Call fmt.Println(bubble([]int{3, 2, 1, 5})) in main to print:

	[1 2 3 5]
	Bonus points: implement Insertion sort
	Extra bonus points: implement an O(n log(n)) sorting algorithm
*/
import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(input []int) []int {
	sorted := make([]int, len(input))
	copy(sorted, input)
	for {
		var swapped = false
		for idx := 0; idx < len(sorted)-1; idx++ {
			if sorted[idx] > sorted[idx+1] {
				swapped = true
				sorted[idx], sorted[idx+1] = sorted[idx+1], sorted[idx]
			}
		}
		if !swapped {
			break
		}
	}
	return sorted
}

func insertion(input []int) []int {
	sorted := make([]int, len(input))
	copy(sorted, input)

	for i := 0; i <= len(sorted)-1; i++ {
		for j := i; j > 0 && sorted[j-1] > sorted[j]; j-- {
			sorted[j], sorted[j-1] =
				sorted[j-1], sorted[j]
		}
	}
	return sorted
}

func merge(input []int) []int {
	// base case. single item lists are inherently sorted
	sorted := make([]int, len(input))
	copy(sorted, input)

	if len(sorted) <= 1 {
		return sorted
	}

	var mid = len(sorted) / 2
	left := sorted[:mid]
	right := sorted[mid:]

	left = merge(left)
	right = merge(right)

	return mergeSlice(left, right)
}

func mergeSlice(left []int, right []int) []int {
	var result = []int{}

	for !(len(left) == 0) && !(len(right) == 0) {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// Either left or right may have elements left; consume them.
	// (Only one of the following loops will actually be entered.)

	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
