package main

import (
	"bytes"
	"strconv"
	"testing"
)

func Test_abs(t *testing.T) {
	tables := []struct {
		nameOfTestcase string
		input          int
		output         int
	}{
		{"normal test 1", 1, 1},
		{"normal test 2", 0, 0},
		{"normal test 3", -1, 1},
	}

	for _, table := range tables {
		actual := abs(table.input)
		expected := table.output
		if expected != actual {
			t.Errorf("Unexpected output from test %s, expected = %v, actual = %v", table.nameOfTestcase, expected, actual)
		}
	}
}

func Test_fib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	tables := []struct {
		nameOfTestcase string
		input          int
		output         string
	}{
		{"normal test 1", 7, strconv.Quote(`1
1
2
3
5
8
13
`)},
		{"normal test 2", 0, strconv.Quote("0\n")},
		{"normal test 3", -7, strconv.Quote(`1
-1
2
-3
5
-8
13
`)},
	}

	for _, table := range tables {
		fib(table.input)
		actual := strconv.Quote(buf.String())
		expected := table.output
		if expected != actual {
			t.Errorf("Unexpected output from test %s, expected = %s, actual = %s", table.nameOfTestcase, expected, actual)
		}

		buf.Reset()
	}
}

func Test_main(t *testing.T) {
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

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
