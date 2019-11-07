package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPAddrString(t *testing.T) {
	var tests = map[string]struct {
		arg  ipAddr
		want string
	}{
		"empty IPv4": {
			ipAddr{},
			"0.0.0.0",
		},
		"8bits IPv4": {
			ipAddr{254},
			"254.0.0.0",
		},
		"16bits IPv4": {
			ipAddr{254, 254},
			"254.254.0.0",
		},
		"24bits IPv4": {
			ipAddr{254, 254, 254},
			"254.254.254.0",
		},
		"full IPv4": {
			ipAddr{254, 254, 254, 254},
			"254.254.254.254",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			actual := test.arg.String()
			assert.Equal(t, test.want, actual)
		})
	}
}

func TestMain(t *testing.T) {
	want := "127.0.0.1\n"
	var buf bytes.Buffer
	out = &buf
	main()
	got := buf.String()
	assert.Equal(t, want, got)
}
