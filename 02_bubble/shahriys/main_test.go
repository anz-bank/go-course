package main

import (
	"bytes"
	"fmt"
	"testing"
)

func testEq(a []int, b []int) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
func TestBubbleSort(t *testing.T) {

	expected := []int{1, 2, 3, 5}
	actual := bubble([]int{3, 2, 1, 5})

	if !testEq(expected, actual) {
		fmt.Print("Expected:")
		fmt.Println(expected)
		fmt.Print("Actual:")
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")

	}

}

func TestInsertionSort(t *testing.T) {

	expected := []int{1, 2, 3, 5}
	actual := insertion([]int{3, 2, 1, 5})

	if !testEq(expected, actual) {
		fmt.Print("Expected:")
		fmt.Println(expected)
		fmt.Print("Actual:")
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")

	}

}
func TestBubbleSortEmptyList(t *testing.T) {

	expected := []int{}
	actual := bubble([]int{})

	if !testEq(expected, actual) {
		fmt.Print("Expected:")
		fmt.Println(expected)
		fmt.Print("Actual:")
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")
	}

}
func TestInsertionSortEmptyList(t *testing.T) {

	expected := []int{}
	actual := insertion([]int{})

	if !testEq(expected, actual) {
		fmt.Print("Expected:")
		fmt.Println(expected)
		fmt.Print("Actual:")
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")
	}

}
func TestBubbleSortMain(t *testing.T) {
	var bufbub bytes.Buffer
	outbub = &bufbub

	main()
	expected := `[1 2 3 5][1 2 3 5]`
	actual := bufbub.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()" + expected + " " + actual)
	}

}
