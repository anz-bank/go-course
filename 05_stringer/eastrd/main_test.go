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

	expected := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestStringer(t *testing.T) {

	expected := "192.168.0.1"
	actual := IPAddr{192, 168, 0, 1}.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestStringer2(t *testing.T) {

	expected := "123.231.0.0"
	actual := IPAddr{123, 231}.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

// Negative inputs and more than 4 integers will both cause overflows
