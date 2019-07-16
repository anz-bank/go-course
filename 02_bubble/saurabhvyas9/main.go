package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	for j := len(r); j > 0; j-- {
		swapped := false
		for i := 0; i < j-1; i++ {
			if r[i] > r[i+1] {
				r[i], r[i+1] = r[i+1], r[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return r
}

func insertion(s []int) []int {
	r := make([]int, len(s))
	copy(r, s)
	for i := 1; i < len(r); i++ {
		temp := r[:i]
		var j int
		for j = 0; j <= len(temp)-1; j++ {
			if temp[j] >= r[i] {
				break
			}
		}
		tmp := r[i]
		copy(r[j+1:i+1], r[j:i])
		r[j] = tmp
	}
	return r
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
}
