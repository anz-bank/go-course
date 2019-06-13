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

	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestNumeronyms(t *testing.T) {
	tables := []struct {
		x []string
		n []string
	}{
		{[]string{"accessibility", "Kubernetes", "abc"}, []string{"a11y", "K8s", "abc"}},
		{[]string{"tres"}, []string{"t2s"}},
		{[]string{"cuatrO"}, []string{"c4O"}},
		{[]string{}, []string{}},
		{[]string{"11111111"}, []string{"161"}},
		{[]string{"?"}, []string{"?"}},
		{[]string{"?$"}, []string{"?$"}},
		{[]string{"?$&%"}, []string{"?2%"}},
	}

	for _, table := range tables {
		numys := numeronyms(table.x...)
		if !equal(numys, table.n) {
			t.Errorf("Sorting of (%v) was incorrect, got: %v, want: %v.", table.x, numys, table.n)
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
