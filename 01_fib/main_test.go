package main

import (
	"strings"
	"bytes"
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
	expected := strconv.Quote("1 1 2 3 5")
	actual := strconv.Quote(strings.TrimSpace(buf.String()))
	r.Equalf(expected, actual, "Unexpected output in main()")
}
