package main

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// MergeSort sorts a given array of integers in ascending order.
func MergeSort(input []int) []int {
	out := make([]int, len(input))
	copy(out, input)

	b := make([]int, len(out))
	n := len(out)

	for width := 1; width < n; width = 2 * width {
		for i := 0; i < n; i += 2 * width {
			merge(out, i, min(i+width, n), min(i+2*width, n), b)
		}
		copyArray(b, out, n)
	}
	return out
}

func merge(a []int, left int, right int, end int, b []int) {
	i := left
	j := right
	for k := left; k < end; k++ {
		if i < right && (j >= end || a[i] <= a[j]) {
			b[k] = a[i]
			i++
		} else {
			b[k] = a[j]
			j++
		}
	}
}

func copyArray(b []int, a []int, n int) {
	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
}
