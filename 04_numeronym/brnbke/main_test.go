package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestLetters(t *testing.T) {
	var testCases = map[string]struct {
		input []string
		want  []string
	}{
		"Empty Strings": {
			input: []string{"", ""},
			want:  []string{"", ""}},
		"Latinish Strings": {
			input: []string{
				"Cat",
				"Donaudampfschiffahrtsgesellschaftskapitän",
				"Hääyöaieuutinen",
				"floccinaucinihilipilification",
				"Speciallægepraksisplanlægningsstabiliseringsperiode",
				"Szczęście",
				"Llanfairpwllgwyngyllgogerychwyrndrobwllllantysiliogogogoch",
			},
			want: []string{"Cat", "D39n", "H13n", "f27n", "S49e", "S7e", "L56h"}},
		"Non Latin Strings": {
			input: []string{
				"⽉⽐⽏⽐⽏⽕⽉⽏⽖",
				"🏦🔫🗯💰🚗😬🚓🚓🚓🚓😱🔫🔫🔫🚓🚓🚒🚑😨😢😰😢😢🚓🏛😢😢😢😢⏸🔒👮",
				"ηλεκτροεγκεφαλογράφημα",
				"فسيكفيكهم"},
			want: []string{"⽉7⽖", "🏦30👮", "η20α", "ف7م"}},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			b := numeronyms(test.input...)
			result := reflect.DeepEqual(test.want, b)
			if !result {
				t.Errorf("want %v, got %v", test.want, b)
			}
		})
	}
}

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("[a11y K8s abc]\n")
	actual := strconv.Quote(buf.String())

	if want != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q \nactual: %q", want, actual)
	}
}
