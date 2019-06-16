package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestFib1Output(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(1)
	expected := strconv.Quote(replaceSpaceWithlineBreak("1 "))
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}

}

func TestFib2Output(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(2)
	expected := strconv.Quote(replaceSpaceWithlineBreak("1 1 "))
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}

}

func TestFib3Output(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(3)
	expected := strconv.Quote(replaceSpaceWithlineBreak("1 1 2 "))
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}

}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(replaceSpaceWithlineBreak("1 1 2 3 5 8 13 "))
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}

}

func replaceSpaceWithlineBreak(line string) string {
	return strings.Replace(line, " ", "\n", -1)
}
