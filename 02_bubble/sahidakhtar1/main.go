package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	list := []int{3, 2, 1, 5}
	bubbleSort(list)
	fmt.Fprintln(out, "After sorting = ", list)
}

func bubbleSort(list []int) {
	count := len(list)
	for count > 0 {
		swappedIndex := 0
		for index := 1; index < count; index++ {
			if list[index-1] > list[index] {
				list[index-1], list[index] = list[index], list[index-1]
				swappedIndex = index
			}
		}
		count = swappedIndex
	}
}
