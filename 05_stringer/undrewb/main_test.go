package main

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("127.0.0.1\n")
	got := strconv.Quote(buf.String())

	assert.Equal(t, got, want)
}

var stringerData = []struct {
	name  string
	input IPAddr
	want  string
}{
	{name: "localhost",
		input: IPAddr{127, 0, 0, 1},
		want:  "127.0.0.1"},
	{name: "four octet addr",
		input: IPAddr{68, 2, 44, 125},
		want:  "68.2.44.125"},
	{name: "three octet addr",
		input: IPAddr{68, 2, 125},
		want:  "68.2.125.0"},
	{name: "two octet addr",
		input: IPAddr{9, 5},
		want:  "9.5.0.0"},
	{name: "one octet addr",
		input: IPAddr{144},
		want:  "144.0.0.0"},
	{name: "empty",
		input: IPAddr{},
		want:  "0.0.0.0"},
}

func TestStringer(t *testing.T) {
	for _, tt := range stringerData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := fmt.Sprint(tt.input)
			assert.Equal(t, got, tt.want)
		})
	}
}
