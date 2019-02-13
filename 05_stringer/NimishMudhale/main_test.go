package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
	expected := fmt.Sprint("127.0.0.1")
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}
func TestDefaultValues(t *testing.T) {
	assert := assert.New(t)
	ip := IPAddr{}
	// When
	actual := ip.String()
	// Then
	expected := fmt.Sprint("0.0.0.0")
	assert.Equal(expected, actual)
}
