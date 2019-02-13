package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests = []struct {
	inputStrings []string
	numeronyms   []string
}{
	{[]string{}, []string{}},
	{[]string{" "}, []string{""}},
	{[]string{"Internationalization", "Localization", "Multilingualization", "accessibility",
		"Canonicalization", "interoperability", "Personalization", "Virtualization"},
		[]string{"I18n", "L10n", "M17n", "a11y", "C14n", "i14y", "P13n", "V12n"}},
	{[]string{" Internationalization ", " Personalization "}, []string{"I18n", "P13n"}},
	{[]string{"Accessibility Localization"}, []string{"A24n"}},
	{[]string{" accessibility localization "}, []string{"a24n"}}}

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestLettersMapAndSort(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		actual := numeronyms(tt.inputStrings...)
		r.EqualValues(tt.numeronyms, actual)
	}
}
