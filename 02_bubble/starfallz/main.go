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
	for i := 0; i < len(s); i++ {
		isSwapped := false
		for j := 0; j < len(s)-1; j++ {
			swap(s, j, &isSwapped)
		}

		if !isSwapped {
			break
		}
	}

	return s
}

func swap(s []int, j int, isSwapped *bool) {
	tempVal := s[j+1]
	if s[j] > tempVal {
		s[j+1] = s[j]
		s[j] = tempVal
		*isSwapped = true
	}
}
