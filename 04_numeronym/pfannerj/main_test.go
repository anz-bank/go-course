package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf
	main()
	expected := "[a11y K8s abc]\n"
	actual := buf.String()
	if expected != actual {
		t.Errorf(expected, actual, "Unexpected output in main()")
	}
}

var numeronymsCases = map[string]struct {
	input, expected []string
}{
	"Standard":         {input: []string{"accessibility", "Kubernetes", "abc"}, expected: []string{"a11y", "K8s", "abc"}},
	"OneTwoThreeFour":  {input: []string{"x", "yz", "xyz", "wxyz"}, expected: []string{"x", "yz", "xyz", "w2z"}},
	"WithSpaceAndPunc": {input: []string{"I like you", "you are very NICE!"}, expected: []string{"I6u", "y12E"}},
	"OneLongArg":       {input: []string{"supercalifragilisticexpialidocious"}, expected: []string{"s32s"}},
	"Return":           {input: []string{"AB\nBA"}, expected: []string{"A2A"}},
	"RawAndStd":        {input: []string{`a11y`, "Kubernetes", `abc`}, expected: []string{"a2y", "K8s", "abc"}},
	"Numeric":          {input: []string{"1234567890"}, expected: []string{"180"}},
	"EscapedQuote":     {input: []string{"a\":2\"b:\"1"}, expected: []string{"a21"}},
	"RawQuote":         {input: []string{`a":2"b:"1`}, expected: []string{"a21"}},
	"EmptyStrings":     {input: []string{``, ""}, expected: []string{"", ""}},
	"RawReturn": {input: []string{`a1
	1y`}, expected: []string{"a2y"}},
}

func TestNumeronyms(t *testing.T) {
	for name, tc := range numeronymsCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := numeronyms(tc.input...)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Errorf("Numeronyms function failed. Input: %v, Actual: %v, Expected: %v", tc.input, actual, tc.expected)
			}
		})
	}
}
