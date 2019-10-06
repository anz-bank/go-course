package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	want := "Read Puppy by ID: {1 Dogo white 50}\n"
	var buf bytes.Buffer
	out = &buf
	main()
	got := buf.String()
	assert.Equal(t, want, got)
}
