package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpaddr(t *testing.T) {
	testData := []struct {
		Scenario string
		input    IPAddr
		expected string
	}{
		{"Happy Day", [4]byte{127, 0, 0, 1}, "127.0.0.1"},
		{"empty", [4]byte{}, "0.0.0.0"},
		{"one byte non zero", [4]byte{0, 0, 0, 10}, "0.0.0.10"},
		{"boundary validation", [4]byte{255, 255, 255, 255}, "255.255.255.255"},
		{"only one byte", [4]byte{10}, "10.0.0.0"},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			actual := td.input.String()
			assert.Equal(t, td.expected, actual)
		})
	}
}
func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "127.0.0.1\n"
	actual := buf.String()
	assert.Equal(expected, actual)
}
