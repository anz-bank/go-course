package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var stdout = os.Stdout

var tests = []struct {
	input    []int
	expected []int
}{
	{[]int{3, 2, 1, 5}, []int{1, 2, 3, 5}},
	{[]int{-71, -99, 89, 4, -82, -73, -65, 94, -92, 99}, []int{-99, -92, -82, -73, -71, -65, 4, 89, 94, 99}},
}

func TestBubble(t *testing.T) {
	for _, test := range tests {
		calculated := bubble(test.input)
		require.Equal(t, test.expected, calculated)

	}
}

func TestInsertion(t *testing.T) {
	for _, test := range tests {
		calculated := insertion(test.input)
		require.Equal(t, test.expected, calculated)

	}
}

func TestMain(t *testing.T) {
	expectedOut := "[1 2 3 5]\n[1 2 3 5]\n"
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	require.Equal(t, expectedOut, result)
}

// captureStart diverts stdio to another file object
func captureStart() (io.Reader, io.Closer) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	return r, w
}

// captureStop copies a file buffer and returns a string of the file
func captureStop(r io.Reader, w io.Closer) string {
	var buf bytes.Buffer
	w.Close()
	os.Stdout = stdout
	_, err := io.Copy(&buf, r)
	if err != nil {
		panic("file not opened")
	}
	return buf.String()
}
