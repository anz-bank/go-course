package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

//main() test
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := `[a11y K8s abc]
`
	actual := buf.String()
	assert.Equalf(t, expected, actual, "got %v want %v", actual, expected)
}

func TestNumeronyms(t *testing.T) {
	testCases := map[string]struct {
		input    []string
		expected []string
	}{
		"test1":  {input: []string{"demon", "GolangCli", "p"}, expected: []string{"d3n", "G7i", "p"}},
		"empty":  {input: []string{}, expected: []string{}},
		"oneval": {input: []string{"a"}, expected: []string{"a"}},
	}
	for name, test := range testCases {
		test := test
		t.Run(name, func(t *testing.T) {
			result := numeronyms(test.input...)
			assert.Equalf(t, test.expected, result, "got %v want %v", result, test.expected)
		})
	}
}
