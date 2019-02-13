package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	list := []int{3, 2, 1, 5}
	sortedList := bubbleSort(list)
	fmt.Fprintln(out, sortedList)
}

func bubbleSort(list []int) []int {
	count := len(list)
	sortedList := make([]int, count)
	copy(sortedList, list)
	for count > 0 {
		swappedIndex := 0
		for index := 1; index < count; index++ {
			if sortedList[index-1] > sortedList[index] {
				sortedList[index-1], sortedList[index] = sortedList[index], sortedList[index-1]
				swappedIndex = index
			}
		}
		count = swappedIndex
	}
	return sortedList
}
