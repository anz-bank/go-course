package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()
	expected := "\"Puppy id: 0\""
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual)
}
