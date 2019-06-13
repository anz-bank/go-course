package main

import "fmt"

func isSorted(s []int) bool {
	for index := range s[1:] {
		if s[index+1] < s[index] {
			return false
		}
	}
	return true
}

func findMin(s []int) int {
	minIndex := 0
	for index, element := range s[1:] {
		if element < s[minIndex] {
			minIndex = index + 1
		}
	}
	return minIndex
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func bubble(s []int) []int {
	copy := s
	sorted := false
	for !sorted {
		for index, element := range copy[1:] {
			if copy[index] > copy[index+1] {
				copy[index+1] = copy[index]
				copy[index] = element
			}
		}
		if isSorted(copy) {
			sorted = true
		}
	}
	return copy
}

func insertion(s []int) []int {
	var sorted []int
	copy := s
	for len(copy) > 0 {
		minIndex := findMin(copy)
		sorted = append(sorted, copy[minIndex])
		copy = remove(copy, minIndex)
	}
	return sorted
}

func quick(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	var lt []int
	var gt []int
	var eq []int
	pivot := s[0]
	for _, element := range s {
		switch {
		case element > pivot:
			gt = append(gt, element)
		case element < pivot:
			lt = append(lt, element)
		default:
			eq = append(eq, element)
		}
	}
	gtSorted := quick(gt)
	ltSorted := quick(lt)
	res := append(ltSorted, eq...)
	res = append(res, gtSorted...)
	return res
}

func main() {
	fmt.Println(bubble([]int{3, 2, 1, 5}))
	fmt.Println(insertion([]int{3, 2, 1, 5}))
	fmt.Println(quick([]int{3, 2, 1, 5}))
}
