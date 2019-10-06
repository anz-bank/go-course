package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeronym(t *testing.T) {
	var tests = map[string]struct {
		arg  string
		want string
	}{
		"empty string":               {"", ""},
		"single string of 1 element": {"a", "a"},
		"single long string":         {"abracadabra", "a9a"},
		"unicode at the start":       {"🤪bracadabra", "🤪9a"},
		"unicode in the middle":      {"abra🤪🤪🤪abra", "a9a"},
		"unicode in the end":         {"abracadabr🤪", "a9🤪"},
		"1 poo":                      {"💩", "💩"},
		"2 poos":                     {"💩💩", "💩💩"},
		"3 poos":                     {"💩💩💩", "💩💩💩"},
		"4 poos":                     {"💩💩💩💩", "💩2💩"},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got := numeronym(test.arg)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNumeronyms(t *testing.T) {
	var tests = map[string]struct {
		arg  []string
		want []string
	}{
		"empty string": {
			[]string{""},
			[]string{""}},
		"mixed strings": {
			[]string{"", "alakazam", "hocuspocus", "mumbojumbo", "💩💩💩💩"},
			[]string{"", "a6m", "h8s", "m8o", "💩2💩"}},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got := numeronyms(test.arg...)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestMain(t *testing.T) {
	want := "[a11y K8s abc]\n"
	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, got, want)
	})
}
