package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringer(t *testing.T) {
	testCases := map[string]struct {
		arg      IPAddr
		expected string
	}{
		"Empty": {
			arg:      IPAddr{},
			expected: "0.0.0.0",
		},
		"Partial": {
			arg:      IPAddr{127},
			expected: "127.0.0.0",
		},
		"LocalHost": {
			arg:      IPAddr{127, 0, 0, 1},
			expected: "127.0.0.1",
		},
		"Common Local": {
			arg:      IPAddr{192, 168, 0, 1},
			expected: "192.168.0.1",
		},
		"Common Local Broadcast": {
			arg:      IPAddr{192, 168, 0, 255},
			expected: "192.168.0.255",
		},
	}

	for testName, tC := range testCases {
		testCase := tC
		t.Run(testName, func(t *testing.T) {
			actualFmt := fmt.Sprint(testCase.arg)
			actualString := testCase.arg.String()

			expected := testCase.expected
			assert.Equal(t, expected, actualFmt)
			assert.Equal(t, expected, actualString)
		})
	}
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	actual := buf.String()
	expected := "127.0.0.1\n"
	assert.Equal(t, expected, actual)
}
