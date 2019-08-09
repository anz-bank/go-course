package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`[a11y K8s abc]`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestNumernym(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  []string
	}{
		"standard": {input: []string{"accessibility", "Kubernetes", "abcd"}, want: []string{"a11y", "K8s", "a2d"}},
		"short":    {input: []string{"qwe", "a", "zx"}, want: []string{"qwe", "a", "zx"}},
		"empty":    {input: []string{"", "", ""}, want: []string{"", "", ""}},
		"mix":      {input: []string{"accessibility", "qwe", ""}, want: []string{"a11y", "qwe", ""}},
		"none":     {input: []string{}, want: []string{}},
		"more": {
			input: []string{"accessibility", "Kubernetes", "abcd", "hiccups", "go", "grub", ""},
			want:  []string{"a11y", "K8s", "a2d", "h5s", "go", "g2b", ""},
		},
	}

	for name, test := range cases {
		input := test.input
		want := test.want
		t.Run(name, func(t *testing.T) {
			assert.Equalf(t, want, numeronyms(input...), "%v was not correct", input)
		})
	}
}
