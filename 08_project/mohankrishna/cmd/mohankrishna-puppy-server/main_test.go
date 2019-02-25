package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLettersMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "{ID:11, Breed:Sheep herder, Colour:Brown, Value:1000}\nNo puppy exists with id 11\n" +
		"{ID:11, Breed:Sheep herder, Colour:Brown, Value:1000}\nNo puppy exists with id 11\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
