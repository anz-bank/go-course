package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	if s == nil || len(s) <= 1 {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1],s[j]
			}
		}
	}

	return s
}

func insertSort(s []int) []int {
	if s == nil || len(s) <= 1 {
		return s
	}

	for i := 1; i < len(s); i++ {
		for j := i - 1; j >= 0; j-- {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1],s[j]
			}else{
				break
			}
		}
	}

	return s
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
