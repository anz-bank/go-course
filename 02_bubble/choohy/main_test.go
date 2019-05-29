package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestBubble(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`Bubble Sort!
[1 2 3 5]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf(expected)
		t.Errorf(actual)
	}
}
