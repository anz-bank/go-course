package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseArgsOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	parseArgs([]string{"--data", "../../test/jsondata.json"})

	// Then
	expected := "{ID:4662, Breed:Sheep herder1, Colour:Brown1, Value:1000}" +
		"{ID:4663, Breed:Sheep herder2, Colour:Brown2, Value:1000}" +
		"{ID:4664, Breed:Sheep herder3, Colour:Brown3, Value:1000}"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestParseArgsOutputWrongFlag(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	parseArgs([]string{"--file", "../../test/jsondata.json"})

	// Then
	expected := "unknown long flag '--file'"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestParseArgsOutputFileDoesntExist(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	parseArgs([]string{"--data", "../../test/data.json"})

	// Then
	expected := "open ../../test/data.json: no such file or directory"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestParseArgsOutputFileCurrupted(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	parseArgs([]string{"--data", "../../test/jsondata_currupted.json"})

	// Then
	expected := "invalid character '{' after array element"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestParseArgsOutputFileDataConflict(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	parseArgs([]string{"--data", "../../test/jsondata.json"})

	// Then
	expected := "A puppy with ID: 4662 already exists"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
	os.RemoveAll("./level_store")
}
