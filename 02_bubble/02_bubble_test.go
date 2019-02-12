package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubbleSortOutput(t *testing.T) {
	// Given
	r := require.New(t)

	// when
	actual := bubble([]int{0, 9, 8, 3, 6, 2})

	//Then
	expected := []int{0, 2, 3, 6, 8, 9}
	r.Equalf(expected, actual, "Unexpected output in main()")

}

func TestBubbleSortOutputWithNegativeElements(t *testing.T) {
	//Given
	r := require.New(t)

	//when
	actual := bubble([]int{-9, -10, -3, -4, -6, -2})

	//Then
	expected := []int{-10, -9, -6, -4, -3, -2}
	r.Equalf(expected, actual, "Unexpected output in main(")
}
