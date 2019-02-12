package main

import (
	"fmt"
)

func main() {
	list := []int{6,3,2,8,1, -1,5}
	fmt.Println("Before sorting = ",list);
	bubbleSort(list)
	fmt.Println("After sorting = ",list);
}

func bubbleSort(list []int) {
	count := len(list)
	swapedIndex := count
	for swapedIndex > 1 {
		for index := 1; index < count ; index ++ {
			if list[index - 1] > list [index] {
				list[index - 1], list[index] = list[index], list[index - 1]
				swapedIndex = index
			}
		}
		count--
	}
}
