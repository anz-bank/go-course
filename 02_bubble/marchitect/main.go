package main

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
)

var out io.Writer = os.Stdout

func bubble(s []int) []int {
	// n is the number of items in the int array
	n := len(s)
	// make a copy of the original int array s
	result := make([]int, n)
	copy(result, s)
	// maybe necessary
	if n < 2 {
		return result
	}
	// set swapped to true
	swapped := true
	// in Go, use keyword for to act as a while loop
	for swapped {
		// set swapped to false
		swapped = false
		// iterate through all of the elements in our list
		for i := 1; i < n; i++ {
			// if the current element is greater than the next element, swap them
			if result[i-1] > result[i] {
				// swap values - Go tuple assignment
				result[i], result[i-1] = result[i-1], result[i]
				// set swap back to true
				swapped = true
			}
		}
	}
	return result
}

func insertion(s []int) []int {
	// n is the number of items in the int array
	n := len(s)
	// make a copy of the original int array s
	result := make([]int, n)
	copy(result, s)
	// maybe necessary
	if n < 2 {
		return result
	}
	var i, j, key int
	for i = 1; i < n; i++ {
		// set the key, a value holder, to be the second item in the int array
		key = result[i]
		// set j to be the item before key in the int array
		j = i - 1
		// for elements of array result[0..i-1], that are greater than key
		for j >= 0 && result[j] > key {
			// sort those greater than key elements one position after current position
			// same as writing result[i] = result[j]
			result[j+1] = result[j]
			j--
		}
		// assign the value of key, or s[i] to the right place
		result[j+1] = key
	}

	return result
}

func quicksort(s []int) []int {
	// n is the number of items in the int array
	n := len(s)
	// maybe necessary
	if n < 2 {
		return s
	}
	// set the boundries: min and max
	left, right := 0, n-1
	// grab a random pivot
	nBig, _ := rand.Int(rand.Reader, big.NewInt(27))
	pivot := int(nBig.Uint64()) % n
	// make the pivot the new max
	s[pivot], s[right] = s[right], s[pivot]
	// iterate through the entire int array
	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}

	s[left], s[right] = s[right], s[left]

	quicksort(s[:left])
	quicksort(s[left+1:])

	return s
}

func main() {
	fmt.Fprint(out, bubble([]int{3, 2, 1, 5, 0, 6, 2, 8}))
}
