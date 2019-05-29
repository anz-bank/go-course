package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestFibo(t *testing.T) {
	tables := []struct {
		x int
		n int
	}{
		{1, 1},
		{7, 13},
		{2, 1},
		{3, 2},
		{6, 8},
		{-6, -8},
		{-3, 2},
		{-4, -3},
	}

	for _, table := range tables {
		total := fib(table.x)
		if total != table.n {
			t.Errorf("Fibo of (%d) was incorrect, got: %d, want: %d.", table.x, total, table.n)
		}
	}
}

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
	13`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}
