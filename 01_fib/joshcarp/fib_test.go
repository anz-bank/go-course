package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var stdout = os.Stdout
var tests = map[string]struct {
	input    int
	expected string
}{
	"positive": {7, "1\n1\n2\n3\n5\n8\n13\n"},
	"negative": {-7, "1\n-1\n2\n-3\n5\n-8\n13\n"},
}

func TestPosFib(t *testing.T) {
	testType := "positive"
	input := tests[testType].input
	r, w := captureStart()
	fib(input)
	result := captureStop(r, w)
	require.Equal(t, result, tests[testType].expected)
}

func TestNegFib(t *testing.T) {
	testType := "negative"
	input := tests[testType].input
	r, w := captureStart()
	fib(input)
	result := captureStop(r, w)
	require.Equal(t, result, tests[testType].expected)
}

func TestMain(t *testing.T) {
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	require.Equal(t, tests["positive"].expected, result)
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
