package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(arr []int) []int {
	r := make([]int, len(arr))
	copy(r, arr)

	for {
		changed := false
		for i := 0; i < len(r)-1; i++ {
			if r[i] > r[i+1] {
				r[i], r[i+1] = r[i+1], r[i]
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	return r
}

func insertion(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)

	for i := 1; i < len(result); i++ {
		tmp := result[i]
		for j := i - 1; j >= 0; j-- {
			if result[j] <= tmp {
				break
			}
			result[j+1], result[j] = result[j], tmp
		}
	}

	return result
}

func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// work out the mid point of array
	i := len(arr) / 2

	// split input arr into two halves and recursively call merge sort to sort each half
	l := mergesort(arr[:i])
	r := mergesort(arr[i:])

	// merge two sorted halves back together.
	return merge(l, r)
}

func merge(l, r []int) []int {

	result := make([]int, 0, len(l)+len(r))

	for len(l) > 0 && len(r) > 0 {
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
	}

	if len(l) > 0 {
		result = append(result, l...)
	} else {
		result = append(result, r...)
	}

	return result
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
