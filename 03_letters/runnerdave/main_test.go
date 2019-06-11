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

	expected := strconv.Quote("aba")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestSortingLetters(t *testing.T) {
	tables := []struct {
		x map[rune]int
		n []string
	}{
		{map[rune]int{97: 1, 96: 3}, []string{"a", "b"}},
	}

	for _, table := range tables {
		sorted := sortLetters(table.x)
		if !equal(sorted, table.n) {
			t.Errorf("Sorting of (%v) was incorrect, got: %v, want: %v.", table.x, sorted, table.n)
		}
	}
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
