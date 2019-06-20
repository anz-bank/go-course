package main

import (
	"bytes"
	"strconv"
	"testing"
)

func Test_main(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`127.0.0.1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}
