package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := "127.0.0.1\n"
	actual := buf.String()

	assert.Equalf(t, expected, actual,
		"Output buffer does not match with expected output.")
}

func TestFmtStringer(t *testing.T) {
	testCases := map[string]struct {
		input       IPAddr
		expectedArr string
	}{
		"127.0.0.1": {input: IPAddr{127, 0, 0, 1},
			expectedArr: "127.0.0.1"},
		"192.168.1.1": {input: IPAddr{192, 168, 1, 1},
			expectedArr: "192.168.1.1"},
		"204.180.10.255": {input: IPAddr{204, 180, 10, 255},
			expectedArr: "204.180.10.255"},
		"0.0.0.0": {input: IPAddr{},
			expectedArr: "0.0.0.0"},
	}

	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, test.input.String(), test.expectedArr,
				"Input does not give expected output.")
		})
	}
}
