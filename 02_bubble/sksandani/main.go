package main

import (
	"fmt"
)

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
	fmt.Println(insertion([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	n := len(s)
	for j := n; j > 0; j-- {
		for i := 1; i <= n-1; i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}
	return s
}

func insertion(s []int) []int {
	n := len(s)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}
	return s
}
