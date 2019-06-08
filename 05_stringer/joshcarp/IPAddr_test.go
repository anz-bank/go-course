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
	input    IPAddr
	expected string
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{255, 255, 255, 255}, "255.255.255.255"},
	{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
}

func TestMain(t *testing.T) {
	expected := "127.0.0.1\n"
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	require.Equal(t, expected, result)
}

func TestStringer(t *testing.T) {
	for _, test := range tests {
		require.Equal(t, test.expected, test.input.String())
	}
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
