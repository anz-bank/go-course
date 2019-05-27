package main

import (
	"bytes"
	"strconv"
	"testing"
	"fmt"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	fib(7)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		fmt.Println(actual)
		t.Errorf("Unexpected output in main()")
	}
}
