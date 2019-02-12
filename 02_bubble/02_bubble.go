package main

import (
	"fmt"
)

func bubble(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				smallerNum := s[j+1]
				s[j+1], s[j] = s[j], smallerNum
			}
		}
	}
	return s
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}
