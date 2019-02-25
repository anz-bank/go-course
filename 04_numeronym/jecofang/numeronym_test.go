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
	// When
	actual := numeronyms("accessibility")

	//Then
	require.Equal(t, "[a11y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsUpperAndLowerCase(t *testing.T) {
	// When
	actual := numeronyms("acceSSibIlitY")

	//Then
	require.Equal(t, "[a11Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumeric(t *testing.T) {
	// When
	actual := numeronyms("abCd12345Y")

	//Then
	require.Equal(t, "[a8Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbols(t *testing.T) {
	// When
	actual := numeronyms("abCd#123%Y")

	//Then
	require.Equal(t, "[a8Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbolsAndUnicode(t *testing.T) {
	// When
	actual := numeronyms("abCd#你好123%Y")

	//Then
	require.Equal(t, "[a10Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsLettersAndNumericAndSymbolsAndUnicodeAndSpace(t *testing.T) {
	// When
	actual := numeronyms("\n \t \r ab \t \r\n Cd  \n #你好123%Y \t")

	//Then
	require.Equal(t, "[a10Y]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsEmpty(t *testing.T) {
	// When
	actual := numeronyms("")

	//Then
	require.Equal(t, "[]", fmt.Sprint(actual), "Unexpected output")
}

func TestNumeronymsSpaces(t *testing.T) {
	// When
	actual := numeronyms("\n\t\r\n   \t \r \n")

	//Then
	require.Equal(t, "[]", fmt.Sprint(actual), "Unexpected output")
}
