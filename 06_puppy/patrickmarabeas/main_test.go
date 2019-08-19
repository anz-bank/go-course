package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

var out io.Writer = os.Stdout

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := ""

	got := buf.String()

	if expected != got {
		t.Errorf("\nExpected: %s\nGot: %s", expected, got)
	}
}
