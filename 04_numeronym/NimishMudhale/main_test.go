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
	expected := fmt.Sprint([]string{"a11y", "K8s", "abc"})
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

func TestEmptyInputs(t *testing.T) {
	// Given
	assert := assert.New(t)
	// When
	numeronymArray := numeronym()
	// Then
	assert.Equal([]string{}, numeronymArray)
}
