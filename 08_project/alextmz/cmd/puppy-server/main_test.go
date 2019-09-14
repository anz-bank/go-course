package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	want := "Read Puppy by ID: 1 {Breed: Dogo, Colour: white, Value 50}.\n"

	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, want, got)
	})
}
