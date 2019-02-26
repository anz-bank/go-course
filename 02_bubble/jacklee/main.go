package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5}))
}

func bubble(origin []int) []int {
	result := make([]int, len(origin))
	copy(result, origin)
	arrLen := len(origin)
	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-1-i; j++ {
			if result[j] > result[j+1] {
				result[j+1], result[j] = result[j], result[j+1]
			}
		}
	}
	return result
}

func heapSort(origin []int) []int {
	heap := make([]int, len(origin))
	copy(heap, origin)
	makeHeap(heap)
	result := make([]int, len(origin))
	heapLen := len(origin)
	for i := 0; i < heapLen; i++ {
		swap(0, len(heap)-1, heap)
		down(0, len(heap)-1, heap)
		n := len(heap)
		x := heap[n-1]
		heap = heap[0 : n-1]
		result[i] = x
	}
	return result
}

func makeHeap(origin []int) {
	n := len(origin)
	for i := n/2 - 1; i >= 0; i-- {
		down(i, n, origin)
	}
}

func down(i, n int, heap []int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		//find less child
		j := j1 //left child
		if j2 := j1 + 1; j2 < n && !less(j1, j2, heap) {
			j = j2 // right child
		}

		if !less(j, i, heap) {
			break
		}
		swap(i, j, heap)
		i = j
	}
}

func less(a, b int, heap []int) bool {
	return heap[a] < heap[b]
}

func swap(a, b int, heap []int) {
	heap[a], heap[b] = heap[b], heap[a]
}
