package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numeronyms(t *testing.T) {
	var tests = map[string]struct {
		arg  []string
		want []string
	}{
		"empty string": {
			[]string{""},
			[]string{""}},
		"single string of 1 element": {
			[]string{"a"},
			[]string{"a"}},
		"single long string": {
			[]string{"abracadabra"},
			[]string{"a9a"}},
		"multiple empty strings": {
			[]string{"", "", ""},
			[]string{"", "", ""}},
		"multiple strings with 1 element": {
			[]string{"a", "b", "c", "d"},
			[]string{"a", "b", "c", "d"}},
		"multiple long strings": {
			[]string{"abracadabra", "alakazam", "hocuspocus", "mumbojumbo"},
			[]string{"a9a", "a6m", "h8s", "m8o"}},
		"mixed strings": {
			[]string{"pirilimpimpim", "MIA", "", "Zzz", "Zoe", "Pizzaz"},
			[]string{"p11m", "MIA", "", "Zzz", "Zoe", "P4z"}},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			got := numeronyms(test.arg...)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_main(t *testing.T) {
	want := "[a11y K8s abc]\n"

	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, got, want)
	})
}
