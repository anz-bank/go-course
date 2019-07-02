package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("&map[]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected result in main() : expected = %s, actual = %s\n", expected, actual)
	}
}
