package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := "127.0.0.1"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}

func TestIPAddr_String(t *testing.T) {
	testCases := map[string]struct {
		input    IPAddr
		expected string
	}{
		"localhost": {input: IPAddr{127, 0, 0, 1}, expected: "127.0.0.1"},
		"gateway":   {input: IPAddr{192, 168, 0, 1}, expected: "192.168.0.1"},
		"google":    {input: IPAddr{172, 217, 25, 142}, expected: "172.217.25.142"},
		"empty":     {input: IPAddr{}, expected: "0.0.0.0"},
		"missing":   {input: IPAddr{192, 163}, expected: "192.163.0.0"},
	}
	for name, test := range testCases {
		input, want := test.input, test.expected
		t.Run(name, func(t *testing.T) {
			actual := fmt.Sprint(input)
			assert.Equal(t, want, actual, "IP address formatted incorrectly")
		})
	}
}
