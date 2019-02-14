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
	expected := "Create Puppy breed = German Shepherd, Colour = white\nUpdate Puppy breed = Pug, Colour = Black\nDelete Puppy breed = , Colour = \n============================================\nPuppy breed = Bulldog, Colour = white\nUpdate Puppy breed = Poodle, Colour = Brown\nDelete Puppy breed = , Colour = \n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
