package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	len := len(s)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if s[j] > s[j+1] {
				temp := s[j]
				s[j] = s[j+1]
				s[j+1] = temp
			}
		}
	}
	return s
}
