package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestTransfer(t *testing.T) {

	testCases := map[string]struct {
		input string
		want  string
	}{
		"normal": {input: "luliangLi", want: "l7i"},
		"blank":  {input: "", want: ""},
		"short":  {input: "lil", want: "lil"},
		"single": {input: "l", want: "l"},
	}

	for name, test := range testCases {
		// t.Run creates a sub test and runs it like a normal test
		test := test

		t.Run(name, func(t *testing.T) {
			actual := transfer(test.input)

			if !reflect.DeepEqual(test.want, actual) {
				t.Errorf("running : %v, expected %v, got %v", test.input, test.want, actual)
			}
		})
	}
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("accessibility", "Kubernetes", "abc")
	expected := []string{"a11y", "K8s", "abc"}
	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected[i], v)
		}
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("running main, expected %v, got %v", expected, actual)
	}
}
