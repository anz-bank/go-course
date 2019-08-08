package main

import (
	"fmt"
)

func bubble(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	var swapped bool
	for {
		swapped = false
		for i := 1; i < len(c); i++ {
			if c[i-1] > c[i] {
				c[i-1], c[i] = c[i], c[i-1]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return c
}

func insertion(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if c[j-1] > c[j] {
				c[j-1], c[j] = c[j], c[j-1]
			}
		}
	}
	return c
}

func mergeSort(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	if len(c) <= 1 {
		return c
	}
	mid := len(c) / 2
	a := mergeSort(c[:mid])
	b := mergeSort(c[mid:])
	return merge(a, b)
}

func merge(a, b []int) []int {
	out := []int{}
	for len(a) > 0 && len(b) > 0 {
		if a[0] <= b[0] {
			out = append(out, a[0])
			a = a[1:]
		} else {
			out = append(out, b[0])
			b = b[1:]
		}
	}

	if len(a) > 0 {
		out = append(out, a...)
	} else {
		out = append(out, b...)
	}
	return out
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}
