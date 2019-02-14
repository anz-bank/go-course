package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuppyMainOutput(t *testing.T) {
	// Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "Create Puppy breed = German Shepherd, Colour = white\n" +
		"Update Puppy breed = Pug, Colour = Black\n" +
		"Delete Puppy breed = , Colour = \n============================================\n" +
		"Puppy breed = Bulldog, Colour = white\n" +
		"Update Puppy breed = Poodle, Colour = Brown\nDelete Puppy breed = , Colour = \n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
