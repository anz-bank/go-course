package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())
	t.Logf("expected:%s", expected)
	t.Logf("actual:%s", actual)

	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestLetters(t *testing.T) {
	tables := []struct {
		n string
		x map[rune]int
	}{
		{"aaa", map[rune]int{'a': 3}},
		{"bbbb", map[rune]int{'b': 4}},
		{"B", map[rune]int{'B': 1}},
		{"bbbB", map[rune]int{'b': 3, 'B': 1}},
		{"bbbBCcccc", map[rune]int{'b': 3, 'B': 1, 'C': 1, 'c': 4}},
		{"bbbBCbbbb", map[rune]int{'b': 7, 'B': 1, 'C': 1}},
		{"", map[rune]int{}},
	}
	for _, table := range tables {
		result := letters(table.n)
		if !reflect.DeepEqual(result, table.x) {
			t.Errorf("Function Letters for (%s) was incorrect, got: %v, want: %v.", table.n, result, table.x)
		}
	}
}

func TestSortingLetters(t *testing.T) {
	tables := []struct {
		x map[rune]int
		n []string
	}{
		{map[rune]int{'a': 1, 'b': 3}, []string{"a:1", "b:3"}},
		{map[rune]int{'c': 1, 'b': 3}, []string{"b:3", "c:1"}},
		{map[rune]int{'m': 1, 'b': 3, 'd': 5}, []string{"b:3", "d:5", "m:1"}},
		{map[rune]int{'a': 1, 'A': 3}, []string{"A:3", "a:1"}},
		{map[rune]int{'n': 1, 'N': 3}, []string{"N:3", "n:1"}},
		{map[rune]int{'n': 0, 'N': 3}, []string{"N:3", "n:0"}},
		{map[rune]int{}, []string{}},
	}

	for _, table := range tables {
		sorted := sortLetters(table.x)
		if !reflect.DeepEqual(sorted, table.n) {
			t.Errorf("Sorting of (%v) was incorrect, got: %v, want: %v.", table.x, sorted, table.n)
		}
	}
}
