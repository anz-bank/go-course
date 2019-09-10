package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IPAddrString(t *testing.T) {
	var tests = map[string]struct {
		arg  IPAddr
		want string
	}{
		"empty IPv4": {
			IPAddr{},
			"0.0.0.0"},
		"8bits IPv4": {
			IPAddr{254},
			"254.0.0.0"},
		"16bits IPv4": {
			IPAddr{254, 254},
			"254.254.0.0"},
		"24bits IPv4": {
			IPAddr{254, 254, 254},
			"254.254.254.0"},
		"full IPv4": {
			IPAddr{254, 254, 254, 254},
			"254.254.254.254"},
	}

	for name, tt := range tests {
		test := tt
		t.Run(name, func(t *testing.T) {
			actual := test.arg.String()
			assert.Equal(t, test.want, actual)
		})
	}
}

func Test_main(t *testing.T) {
	want := "127.0.0.1\n"

	t.Run("main test", func(t *testing.T) {
		var buf bytes.Buffer
		out = &buf
		main()
		got := buf.String()
		assert.Equal(t, got, want)
	})
}
