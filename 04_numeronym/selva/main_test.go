package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var nemeronymsTest = []struct {
	in  []string
	out []string
}{
	{[]string{"accessibility", "internationalisation"}, []string{"a11y", "i18n"}},
}

var numeronymsTest = []struct {
	in          []string
	expectedOut []string
}{
	{[]string{"accessibility", "internationalisation"}, []string{"a11y", "i18n"}},
	{[]string{}, []string{}},
	{[]string{"localisation    ", "  "}, []string{"l10n", ""}},
}

func TestNumeronyms(t *testing.T) {
	r := require.New(t)
	for _, tt := range numeronymsTest {
		out := numeronyms(tt.in...)
		r.EqualValues(tt.expectedOut, out, "Unexpected output in main()")
	}
}
