package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	fibonnaciArray := []int{1, 1, 2, 3, 5, 8, 13}
	expected := fmt.Sprint("Fibonnaci Series of 7:", fibonnaciArray)
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
