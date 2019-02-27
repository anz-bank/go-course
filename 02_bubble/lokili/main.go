package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(arr []int) []int {
	arrLen := len(arr)
	sorted := make([]int, arrLen)
	copy(sorted, arr)

	for i := 0; i < arrLen-1; i++ {
		swapped := false
		for j := 0; j < arrLen-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return sorted
}

func insertion(arr []int) []int {
	arrLen := len(arr)
	sorted := make([]int, arrLen)
	copy(sorted, arr)

	for i := 1; i <= arrLen-1; i++ {
		insertValue := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > insertValue {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = insertValue
	}
	return sorted
}

func mergeSort(arr []int, start int, mid int, end int, tmp []int) {
	i := start
	j := mid + 1
	k := 0

	for i <= mid && j <= end {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}
	for j <= end {
		tmp[k] = arr[j]
		k++
		j++
	}

	k = 0
	for start <= end {
		arr[start] = tmp[k]
		start++
		k++
	}
}

func mergeSortUp2Down(arr []int, start int, end int, tmp []int) {
	if arr == nil || start >= end {
		return
	}

	mid := (start + end) / 2
	mergeSortUp2Down(arr, start, mid, tmp)
	mergeSortUp2Down(arr, mid+1, end, tmp)

	mergeSort(arr, start, mid, end, tmp)
}

func merge(arr []int) []int {
	arrLen := len(arr)
	sorted := make([]int, arrLen)
	copy(sorted, arr)
	tmp := make([]int, arrLen)
	mergeSortUp2Down(sorted, 0, arrLen-1, tmp)

	return sorted
}

func main() {
	unsorted := []int{3, 1, 2, 5}
	fmt.Fprint(out, bubble(unsorted))
}
