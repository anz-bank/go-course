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
	input    []string
	expected []string
}{
	{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
	{[]string{"summary", "dive", "vat", "a", "stunning"}, []string{"s5y", "d2e", "vat", "a", "s6g"}},
}

func TestMain(t *testing.T) {
	expected := "[a11y K8s abc]\n"
	r, w := captureStart()
	main()
	result := captureStop(r, w)
	require.Equal(t, expected, result)
}

func TestNeumonym(t *testing.T) {
	for _, test := range tests {
		require.Equal(t, test.expected, numeronyms(test.input...))
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
