package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`1
1
2
3
5
8
13
`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibo(t *testing.T) {
	tables := []struct {
		x int
		n string
	}{
		{0, "0 "},
		{1, "1 "},
		{7, "1 1 2 3 5 8 13 "},
		{2, "1 1 "},
		{3, "1 1 2 "},
		{6, "1 1 2 3 5 8 "},
		{-6, "1 -1 2 -3 5 -8 "},
		{-3, "1 -1 2 "},
		{-4, "1 -1 2 -3 "},
	}

	for _, table := range tables {
		var buf bytes.Buffer
		out = &buf
		fib(table.x)
		expected := strconv.Quote(table.n)
		actual := strconv.Quote(strings.Replace(buf.String(), "\n", " ", -1))
		if expected != actual {
			t.Errorf("Fibo of (%d) was incorrect, got: %s, want: %s.", table.x, actual, expected)
		}
	}
}
