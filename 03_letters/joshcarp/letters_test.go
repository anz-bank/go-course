package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	expected := "a:2\nb:1\n"
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	require.Equal(t, expected, result)
}

// captureStart diverts stdio to another file object
func captureStart() (*os.File, *os.File) {
	r, w, _ := os.Pipe()
	os.Stdout = w
	return r, w
}

// captureStop copies a file buffer and returns a string of the file
func captureStop(r *os.File, w *os.File) string {
	var buf bytes.Buffer
	w.Close()
	io.Copy(&buf, r)
	return buf.String()
}
