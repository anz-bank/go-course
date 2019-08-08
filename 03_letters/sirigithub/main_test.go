package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterFrequency(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    []string
	}{
		{description: "Simple string", input: "thisisasimplestring",
			expected: []string{"a:1", "e:1", "g:1", "h:1", "i:4", "l:1", "m:1", "n:1", "p:1", "r:1", "s:4", "t:2"}},

		{description: "String with spaces", input: "   223  ", expected: []string{" :5", "2:2", "3:1"}},

		{description: "Empty String", input: "", expected: []string{}},

		{description: "German Umlauts", input: "äöß€’üüüöäßß", expected: []string{"ß:3", "ä:2", "ö:2", "ü:3", "’:1", "€:1"}},

		{description: "String with Back slashes", input: "a\\c\\b", expected: []string{"\\:2", "a:1", "b:1", "c:1"}},

		{description: "String with Emojis", input: "😄🐷🙈🐷🏃😄", expected: []string{"🏃:1", "🐷:2", "😄:2", "🙈:1"}},
	}

	for _, test := range tests {
		actual := sortLetters(letters(test.input))
		expected := test.expected
		t.Run(test.description, func(t *testing.T) {
			assert.Equal(t, actual, expected, "actual %v but expected %v", actual, expected)
		})
	}
}

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
