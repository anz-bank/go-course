package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "a:2\nb:1\n"
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output in main(). Expected = %v Actual = %v", expected, actual)
	}
}

func TestLetterFrequency(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "Empty input", input: "", expected: ""},
		{name: "Lower and upper cases", input: "aAbBcC", expected: "A:1 B:1 C:1 a:1 b:1 c:1"},
		{name: "Numbers", input: "98765432109876543210", expected: "0:2 1:2 2:2 3:2 4:2 5:2 6:2 7:2 8:2 9:2"},
		{name: "Normal ASCII inputs", input: "abcabc123123", expected: "1:2 2:2 3:2 a:2 b:2 c:2"},
		{name: "Special Character inputs", input: "!@#$%^&*!@#$%^&*", expected: "!:2 #:2 $:2 %:2 &:2 *:2 @:2 ^:2"},
		{name: "Hello World in Chinese", input: "世界您好", expected: "世:1 好:1 您:1 界:1"},
		{name: "Emojis", input: "\u263A\u263E", expected: "\u263A:1 \u263E:1"},
	}

	for _, test := range tests {
		//replace spaces with newline character to match expected output - easier to read expected output
		test.expected = strings.Replace(test.expected, " ", "\n", -1)
		actual := strings.Join(sortLetters(letters(test.input)), "\n")
		assert.Equal(t, test.expected, actual, "Test Failed!")
	}
}
