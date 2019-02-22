package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	r.Equalf(`[a11y K8s abc]`, buf.String(), "Unexpected output in main()")
}

func TestNumeronymsLettersOnly(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("accessibility")

	//Then
	r.Equalf("[a11y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsUpperAndLowerCase(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("acceSSibIlitY")

	//Then
	r.Equalf("[a11Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumeric(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("abCd12345Y")

	//Then
	r.Equalf("[a8Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbols(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("abCd#123%Y")

	//Then
	r.Equalf("[a8Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbolsAndUnicode(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("abCd#你好123%Y")

	//Then
	r.Equalf("[a10Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbolsAndUnicodeAndSpace(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("\n \t \r ab \t \r\n Cd  \n #你好123%Y \t")

	//Then
	r.Equalf("[a10Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsEmpty(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("")

	//Then
	r.Equalf("[]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsSpaces(t *testing.T) {
	// Given
	r := require.New(t)

	// When
	actual := numeronyms("\n\t\r\n   \t \r \n")

	//Then
	r.Equalf("[]", fmt.Sprint(actual), "Unexpected output")
}
