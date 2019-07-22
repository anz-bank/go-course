package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := "\"Puppy id: 0\""
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output\nexpected: %v\nactual: %v", expected, actual)
	}
}
