package main

import (
	"bytes"
	"fmt"
	"strconv"
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
	expected := strconv.Quote("\n1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestNegativeFib(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	fmt.Fprintln(out, fib(-7))

	// Then
	expected := strconv.Quote("\n1\n-1\n2\n-3\n5\n-8\n13\n")
	actual := strconv.Quote(buf.String())
	r.Equalf(expected, actual, "Unexpected output in main()")

}
