package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIpaddr(t *testing.T) {
	testData := []struct {
		Scenario string
		input    IPAddr
		expected string
	}{
		{"Scenario One", [4]byte{127, 0, 0, 1}, "127.0.0.1"},
		{"Scenario two", [4]byte{}, "0.0.0.0"},
		{"Scenario three", [4]byte{0, 0, 0, 10}, "0.0.0.10"},
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
	expected := "127.0.0.1"
	actual := buf.String()
	actual = strings.Replace(actual, "\n", ",", -1)
	actual = strings.TrimRight(actual, ",")
	assert.Equal(expected, actual)
}
