package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestIPAddrStringer(t *testing.T) {
	var testCases = map[string]struct {
		input [4]byte
		want  string
	}{
		"Random IP": {
			input: [4]byte{217, 160, 230, 133},
			want:  "217.160.230.133"},
		"zeros IP": {
			input: [4]byte{0, 0, 0, 0},
			want:  "0.0.0.0"},
		"Empty IP": {
			input: [4]byte{},
			want:  "0.0.0.0"},
		"Partial IP": {
			input: [4]byte{5, 4, 5},
			want:  "5.4.5.0"},
	}

	for name, test := range testCases {
		test := test

		t.Run(name, func(t *testing.T) {
			ip := IPAddr(test.input)
			if test.want != ip.String() {
				t.Errorf("want %v, got %v", test.want, ip)
			}
		})
	}
}
func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	want := strconv.Quote("127.0.0.1\n")
	actual := strconv.Quote(buf.String())

	if want != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", want, actual)
	}
}
