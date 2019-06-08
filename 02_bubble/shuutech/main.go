package main

import (
	"fmt"
)

func bubble(s []int) []int {
	n := len(s)
	var swapped bool
	for ok := true; ok; ok = swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				a := s[i-1]
				b := s[i]
				s[i-1] = b
				s[i] = a
				swapped = true
			}
		}
	}
	return s
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}
