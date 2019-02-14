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
	{[]string{""}, []string{""}},
	{[]string{"Internationalization", "Localization", "Multilingualization", "accessibility",
		"Canonicalization", "interoperability", "Personalization", "Virtualization"},
		[]string{"I18n", "L10n", "M17n", "a11y", "C14n", "i14y", "P13n", "V12n"}},
	{[]string{"Internationalization ", " Personalization"}, []string{"I19 ", " 14n"}},
	{[]string{"Accessibility Localization"}, []string{"A24n"}},
	{[]string{" accessibility localization "}, []string{" 26 "}}}

func TestNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, tt := range tests {
		actual := numeronyms(tt.inputStrings...)
		r.EqualValues(tt.numeronyms, actual)
	}
}
func TestNumeroniumMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "[a11y K8s abc]\n"

	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestNumeroniumsNoArgs(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := numeronyms()

	// Then
	expected := 0
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestNumeroniumsSingleArg(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := numeronyms("ahajka")

	// Then
	expected := 1
	actual := len(m)
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsEmptyString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("")

	// Then
	expected := ""
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsSingleCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("a")

	// Then
	expected := "a"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestGetNumeroniumsDoubleCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("ab")

	// Then
	expected := "ab"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsThreeCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("abc")

	// Then
	expected := "abc"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestGetNumeroniumsFourCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("abcd")

	// Then
	expected := "a2d"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestGetNumeroniumsMultiCharString(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	m := getNumeronym("Global Warming")

	// Then
	expected := "G12g"
	actual := m
	r.Equalf(expected, actual, "Unexpected output in main()")
}
