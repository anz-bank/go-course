package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func numeroniumTest(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]\n"

	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func testNoArgument(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := numeronyms()

	// Then
	expected := 0
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func testNumeroniumsSingleArg(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	a := numeronyms("abc")

	// Then
	expected := 1
	actual := len(a)
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func emptyString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := prepareResponse("")

	// Then
	expected := ""
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func lessThanThreeCharTesting(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := prepareResponse("ab")

	// Then
	expected := "ab"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func threeCharTesting(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := prepareResponse("abc")

	// Then
	expected := "abc"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func testGetNumeroniumsMultiCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := prepareResponse("Good morning")

	// Then
	expected := "G10g"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
