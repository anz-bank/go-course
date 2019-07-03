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
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestStringer(t *testing.T) {
	tables := []struct {
		input  IPAddr
		output string
	}{
		{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{IPAddr{10, 0, 0, 1}, "10.0.0.1"},
		{IPAddr{191, 255, 255, 255}, "191.255.255.255"},
		{IPAddr{10, 10}, "10.10.0.0"},
		{IPAddr{}, "0.0.0.0"},
	}

	for _, table := range tables {
		expected := table.output
		actual := table.input.String()
		if expected != actual {
			t.Errorf("Stringer of (%#v) was incorrect, got: %s, want: %s.", table.input, actual, expected)
		}
	}
}
