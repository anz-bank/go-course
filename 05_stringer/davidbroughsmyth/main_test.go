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
	expected := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())
	assert.Equal(t, expected, actual, "Unexpected output from main()")
}

var tests = map[string]struct {
	input IPAddr
	want  string
}{
	"Empty":     {input: IPAddr{}, want: "0.0.0.0"},
	"Loopback":  {input: IPAddr{127, 0, 0, 1}, want: "127.0.0.1"},
	"Partial":   {input: IPAddr{0, 255}, want: "0.255.0.0"},
	"Broadcast": {input: IPAddr{255, 255, 255, 255}, want: "255.255.255.255"},
	"Zero mask": {input: IPAddr{0, 0, 0, 0}, want: "0.0.0.0"},
	"Local lan": {input: IPAddr{192, 168, 0, 1}, want: "192.168.0.1"},
}

func TestIPAddr(t *testing.T) {
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			got := test.input.String()
			assert.Equal(t, test.want, got)
		})
	}
}
