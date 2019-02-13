package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInputs = []struct {
	in []string
	out []string 
}{
	{[]string{`1234`, `32232`, `#@$@#$@`, "a", "accessibility", ""},[]string{`124`, `332`, `#5@`, "a", "a11y", ""}}}
	


func TestNumronyms(t *testing.T) {
	// Given
	r := require.New(t)
	for _, t := range testInputs {
		actual := numeronyms(t.in...)
		r.EqualValues(actual, t.out)
	}
}	