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

	sCopy := append(s[:0:0], s...)
	for i := 0; i < len(sCopy)-1; i++ {
		for j := 0; j < len(sCopy)-1-i; j++ {
			if sCopy[j] > sCopy[j+1] {
				sCopy[j], sCopy[j+1] = sCopy[j+1], sCopy[j]
			}
		}
	}

	return sCopy
}

func insertSort(s []int) []int {
	if s == nil || len(s) <= 1 {
		return s
	}

	sCopy := append(s[:0:0], s...)
	for i := 1; i < len(sCopy); i++ {
		for j := i - 1; j >= 0; j-- {
			if sCopy[j] > sCopy[j+1] {
				sCopy[j], sCopy[j+1] = sCopy[j+1], sCopy[j]
			} else {
				break
			}
		}
	}

	return sCopy
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
}
