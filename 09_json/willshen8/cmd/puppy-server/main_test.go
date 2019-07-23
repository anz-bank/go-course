package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	args = []string{"--wrongFlag"}
	assert.Panics(t, main)
}
func TestLongFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	args = []string{"--data", "puppies.json"}

	main()
	expected := `{1 Jack Russell Terrier White and Brown 1500}
{1234 Fox Terrier Black 1300}
{100 German Shepperd Brown 2000}
{120 Golden Retriever Golden 2500}
{200 Chihuahua White 500}
{300 Husky White 3500}
{700 Pomeranian White 700}
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestShortFlag(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	args = []string{"-d", "puppies.json"}

	main()
	expected := `{1 Jack Russell Terrier White and Brown 1500}
{1234 Fox Terrier Black 1300}
{100 German Shepperd Brown 2000}
{120 Golden Retriever Golden 2500}
{200 Chihuahua White 500}
{300 Husky White 3500}
{700 Pomeranian White 700}
`
	actual := buf.String()
	assert.Equal(t, expected, actual)
}

func TestBadFileName(t *testing.T) {
	args = []string{"--data=helloWorld.json"}
	assert.Panics(t, main)
}

func TestSavePuppiesToStore(t *testing.T) {
	args = []string{"--data=invalid_puppies.json"}
	assert.Panics(t, main)
}

func TestUnmarshalJSON(t *testing.T) {
	args = []string{"--data=type_mismatch.json"}
	assert.Panics(t, main)
}
