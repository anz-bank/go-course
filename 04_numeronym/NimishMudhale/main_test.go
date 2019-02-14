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

var cases = []struct {
	in, out []string
}{
	{[]string{"internationalization", "Administration", "consumability"}, []string{"i18n", "A12n", "c11y"}},
	{[]string{}, []string{}},
	{[]string{""}, []string{""}},
	{[]string{"🦄☣🦄☣"}, []string{"🦄2☣"}},
}

func TestInputs(t *testing.T) {
	//Given
	assert := assert.New(t)
	for _, e := range cases {
		res := numeronym(e.in...)
		assert.Equal(e.out, res)
	}
}
