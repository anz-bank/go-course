package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(items []int) []int {

	if len(items) > 1 {
		pivotIndex := len(items) / 2
		var smallerItems = []int{}
		var largerItems = []int{}

		for i := range items {
			val := items[i]
			if i != pivotIndex {
				if val < items[pivotIndex] {
					smallerItems = append(smallerItems, val)
				} else {
					largerItems = append(largerItems, val)
				}
			}
		}

		bubble(smallerItems)
		bubble(largerItems)

		var merged []int = append(append(append([]int{}, smallerItems...), []int{items[pivotIndex]}...), largerItems...)

		for j := 0; j < len(items); j++ {
			items[j] = merged[j]
		}

	}

	return items

}

func main() {
	a := []int{3, 2, 1, 5}
	fmt.Fprintln(out, bubble(a))
}
