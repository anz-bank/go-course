package main

import (
	"fmt"
)

func bubble(s []int) []int {
	var n = len(s)
	for j := 0; j < n; j++ {
		for i := 0; i < n - 1; i++ {
			if s[i] > s[i + 1] {
				s[i], s[i + 1] = s[i + 1], s[i]
			}
		}
	}
	return s
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}