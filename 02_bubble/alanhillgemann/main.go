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

func bubble(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	for i := len(sortedArr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if sortedArr[j] > sortedArr[j+1] {
				sortedArr[j], sortedArr[j+1] = sortedArr[j+1], sortedArr[j]
			}
		}
	}
	return sortedArr
}
