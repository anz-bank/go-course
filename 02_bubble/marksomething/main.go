package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	// Copy slice: See https://github.com/go101/go101/wiki
	sorted := append(s[:0:0], s...)

	for range sorted {
		noSwaps := true

		// need to do n-1 compare/swaps
		for i := 0; i < len(s)-1; i++ {
			if sorted[i] > sorted[i+1] {
				sorted[i], sorted[i+1] = sorted[i+1], sorted[i]
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}

	return sorted
}

func insertion(s []int) []int {
	//pre-allocate output, maintain high water mark in loop
	sorted := make([]int, len(s))

	for hwm, needle := range s {

		//find insertion point
		insPoint := hwm
		for i, v := range sorted[:hwm] {
			if needle < v {
				insPoint = i
				break
			}
		}

		//move data at and after insertion point up by 1 and insert new value
		copy(sorted[insPoint+1:hwm+1], sorted[insPoint:hwm])
		sorted[insPoint] = needle
	}
	return sorted
}

func quick(s []int) []int {
	// Quicksorting inplace, so make a copy
	// Copy slice: See https://github.com/go101/go101/wiki
	sorted := append(s[:0:0], s...)

	// Need to predefine so we can recurse
	var qsort func([]int, int, int)
	qsort = func(slc []int, low int, high int) {
		// no value or single value are trivially sorted
		if low == high || low+1 == high {
			return
		}

		pVal := slc[high-1]
		pLoc := low

		//swap all values less than pivot value
		//with pivot location and increment pivot
		for i := range slc[low : high-1] {
			if slc[low+i] < pVal {
				slc[low+i], slc[pLoc] = slc[pLoc], slc[low+i]
				pLoc++
			}
		}
		slc[high-1], slc[pLoc] = slc[pLoc], slc[high-1]

		qsort(slc, low, pLoc)
		qsort(slc, pLoc+1, high)
	}
	qsort(sorted, 0, len(sorted))
	return sorted
}

func main() {
	fmt.Fprintln(out, bubble([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, insertion([]int{3, 2, 1, 5}))
	fmt.Fprintln(out, quick([]int{3, 2, 1, 5}))
}
