package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

// this testing function aims to test two functions letters() and sortLetters() from main, together
func Test_letters_and_sortLetters(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          string
		output         []string
	}{
		{"normal test 1", "test", []string{"e:1", "s:1", "t:2"}},
		{"normal test 2", "$Bfs`BY", []string{"$:1", "B:2", "Y:1", "`:1", "f:1", "s:1"}},
		{"empty input", "", []string{}},
		{"Mandarin Characters Hello", "你好", []string{"你:1", "好:1"}},
		{"normal test 3", "Brain drained", []string{" :1", "B:1", "a:2", "d:2", "e:1", "i:2", "n:2", "r:2"}},
	}
	for _, table := range tables {
		expected := sortLetters(letters(table.input))
		actual := table.output
		// reference: https://yourbasic.org/golang/compare-slices/
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Unexpected output from test %s, expected = %s, actual = %s", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_main(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`a:2
b:1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}
