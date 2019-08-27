package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := "Puppy id: 0"
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output\nexpected: %v\nactual: %v", expected, actual)
	}
}
