package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "hello"

	got := buf.String()

	if expected != got {
		t.Errorf("\nExpected: %s\nGot: %s", expected, got)
	}
}
