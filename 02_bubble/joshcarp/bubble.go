package main

import "fmt"

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}
func bubble(s []int) []int {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
