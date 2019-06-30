package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("[a11y K8s abc]\n")
	got := strconv.Quote(buf.String())

	if got != want {
		t.Errorf("actual: %s does not match expected: %s", got, want)
	}
}

var numeronymsData = []struct {
	name  string
	input []string
	want  []string
}{
	{name: "lab example",
		input: []string{"accessibility", "Kubernetes", "abc"},
		want:  []string{"a11y", "K8s", "abc"}},
	{name: "empty",
		input: []string{},
		want:  []string{}},
}

func TestNumeronyms(t *testing.T) {
	for _, tt := range numeronymsData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := numeronyms(tt.input...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numeronyms() = \n%v,\nwant \n%v", got, tt.want)
			}
		})
	}
}

var numeronymData = []struct {
	name  string
	input string
	want  string
}{
	{name: "lab example", input: "Kubernetes", want: "K8s"},
	{name: "trailing newline", input: "aba", want: "aba"},
	{name: "duplicate entries", input: "KuKuKuKuK", want: "K7K"},
	{name: "emptyr", input: "", want: ""},
	{name: "space example", input: "blah blah", want: "b7h"},
	{name: "i1bn", input: "internationalization", want: "i18n"},
}

func TestNumeronym(t *testing.T) {
	for _, tt := range numeronymData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := numeronym(tt.input)
			if got != tt.want {
				t.Errorf("numeronym() = \n%v,\nwant \n%v", got, tt.want)
			}
		})
	}
}
