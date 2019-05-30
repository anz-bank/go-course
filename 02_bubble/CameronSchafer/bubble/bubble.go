package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	check := false

	for !check {
		tempCheck := true
		//sort once and then check
		for i := 0; i < len(s); i++ {
			//check if in last slice
			if i+1 < len(s) {
				if s[i] > s[i+1] {
					//swap the elems around
					s[i], s[i+1] = s[i+1], s[i]
					tempCheck = false
				}
			}
		}
		check = tempCheck
	}
	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
