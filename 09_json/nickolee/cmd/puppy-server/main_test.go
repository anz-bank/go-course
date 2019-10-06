package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	args = []string{"-d", "../../puppy-data/puppies.json"}

	main()

	expected := `Puppy with ID 1 has been created
Retrieved puppy: &{1 Vulpix Red 2900}
Puppy with ID 2 has been created
Retrieved puppy: &{2 Eevee Light Brown 1290}
Puppy with ID 3 has been created
Retrieved puppy: &{3 Vaporeon Sea Blue 3290}
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestLongFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	args = []string{"--data", "../../puppy-data/puppies.json"}

	main()

	expected := `Puppy with ID 1 has been created
Retrieved puppy: &{1 Vulpix Red 2900}
Puppy with ID 2 has been created
Retrieved puppy: &{2 Eevee Light Brown 1290}
Puppy with ID 3 has been created
Retrieved puppy: &{3 Vaporeon Sea Blue 3290}
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestParseError(t *testing.T) {
	args = []string{"--wrongFlag"}
	assert.Panics(t, main)
}

func TestEmptyFileName(t *testing.T) {
	args = []string{"--data"}
	assert.Panics(t, main)
}

func TestWrongFileName(t *testing.T) {
	args = []string{"--data", "iDontExist.json"}
	assert.Panics(t, main)
}

func TestUnmarshalPuppiesTypeMismatch(t *testing.T) {
	args = []string{"--data", "../../puppy-data/type_mismatch.json"}
	assert.Panics(t, main)
}

func TestUnmarshalPuppiesInvalidJSON(t *testing.T) {
	args = []string{"--data", "../../puppy-data/invalid.json"}
	assert.Panics(t, main)
}
