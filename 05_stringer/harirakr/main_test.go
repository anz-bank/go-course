package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "127.0.0.1\n", buf.String())
}

func TestStringer(t *testing.T) {
	tcs := map[string]struct {
		want string
		ip   IPAddr
	}{
		"null":   {want: "0.0.0.0", ip: IPAddr{}},
		"zero":   {want: "0.0.0.0", ip: IPAddr{0}},
		"classA": {want: "10.0.0.0", ip: IPAddr{10}},
		"classB": {want: "192.168.0.0", ip: IPAddr{192, 168}},
		"classC": {want: "172.16.100.0", ip: IPAddr{172, 16, 100}},
		"classD": {want: "8.8.8.8", ip: IPAddr{8, 8, 8, 8}},
		"mask":   {want: "255.255.255.255", ip: IPAddr{255, 255, 255, 255}},
	}

	for name, tc := range tcs {
		output := tc.ip.String()
		want := tc.want
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, want, output)
		})
	}
}
