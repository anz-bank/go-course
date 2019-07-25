package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	var b bytes.Buffer
	out = &b
	expected := format("1 1 2 3 5 8 13 ")
	main()
	actual := b.String()
	fmt.Print(actual)
	if expected != actual {
		t.Errorf("Results are not matching exptected %v,actual %v", expected, actual)
	}
}
func TestFib(t *testing.T) {
	var b bytes.Buffer
	out = &b
	testRecords := []struct {
		n        int
		expected string
	}{
		{6, format("1 1 2 3 5 8 ")},
		{-7, format("1 -1 2 -3 5 -8 13 ")},
		{0, ""},
		{2, format("1 1 ")},
	}
	for _, testRec := range testRecords {
		fib(testRec.n)
		actual := b.String()
		if actual != testRec.expected {
			t.Errorf("Results are not matching exptected %v,actual %v", testRec.expected, actual)
		}
		b.Reset()
	}
}
func format(data string) string {
	return strings.ReplaceAll(data, " ", "\n")
}
