package main

import (
	"bytes"
	"strconv"
	"testing"
)

func Test_main(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`127.0.0.1
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main(), expected = %v, actual = %v", expected, actual)
	}
}

func Test_String(t *testing.T) {
	testCases := map[string]struct {
		input  IPAddr
		output string
	}{
		"Empty": {
			input:  IPAddr{},
			output: "0.0.0.0",
		},
		"Partially Implemented": {
			input:  IPAddr{1, 2},
			output: "1.2.0.0",
		},
		"Example IP Address": {
			input:  IPAddr{172, 16, 80, 100},
			output: "172.16.80.100",
		},
		"Broadcast": {
			input:  IPAddr{255, 255, 255, 255},
			output: "255.255.255.255",
		},
		"Localhost": {
			input:  IPAddr{127, 0, 0, 1},
			output: "127.0.0.1",
		},
	}

	for testCase, test := range testCases {
		input, expected := test.input, test.output
		actual := input.String()
		if actual != expected {
			t.Errorf("Unexpected output from test %s, expected = %s, actual = %s", testCase, expected, actual)
		}
	}

}
