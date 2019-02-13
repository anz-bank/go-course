package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputs = []struct {
	in []string
	out []string 
}{
	{`1234`, `32232`},{`124`, `332`}}


func TestNumronyms(t *testing.T) {
	// Given
	r := require.New(t)
	for _, t := range testInputs {
		actual := numeronyms(t.in)
		r.EqualValues(actual, out)
	}
}	