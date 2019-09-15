package main

import (
	"bytes"
	"testing"
)

type controlTest struct {
	input    IPAddr
	expected string
}

var hosts = []controlTest{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	{IPAddr{}, "0.0.0.0"},
}

func TestIPs(t *testing.T) {

	for _, tt := range hosts {
		actual := tt.input.String()
		if actual != tt.expected {
			t.Errorf("ip stringer(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}

func TestMain(t *testing.T) {
	expected := "loopback: 127.0.0.1\ngoogleDNS: 8.8.8.8\n"
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()

	if actual != expected {
		t.Errorf("main expected %v, got %v", expected, actual)
	}
}
