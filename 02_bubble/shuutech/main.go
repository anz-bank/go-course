package main

import (
	"fmt"
	"sort"
)

func bubble(s []int) []int {
	sort.Ints(s)

	return s
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}
