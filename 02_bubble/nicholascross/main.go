package main

import "fmt"

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {
	count := len(s)
	sort := s[:]

	for {
		swapped := false

		for i := 0; i+1 < count; i++ {
			if sort[i] > sort[i+1] {
				sort[i+1], sort[i] = sort[i], sort[i+1]
				swapped = true
			}
		}

		if !swapped {
			//When nothing is swapped it means the order of all subsequent
			//elements is ascending and we can start bubbling up the next value
			break
		}
	}

	return sort
}
