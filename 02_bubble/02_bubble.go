package main

import (
	"fmt"
)

func main() {
	fmt.Println(bubble([]int{3, 4, 1, 5, 2}))
}

func bubble(s []int) []int {
	var temp = 0
	for i := 1; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if s[j] > s[j+1] {
				temp = s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	return s
}
