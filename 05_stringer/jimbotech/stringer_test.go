package main

import (
	"bytes"
	"testing"
)

func TestIPs(t *testing.T) {
	type controlTest struct {
		input    ipAddr
		expected string
	}
	var hosts = []controlTest{
		{ipAddr{127, 0, 0, 1}, "127.0.0.1"},
		{ipAddr{8, 8, 8, 8}, "8.8.8.8"},
		{ipAddr{}, "0.0.0.0"},
		{ipAddr{10}, "10.0.0.0"},
		{ipAddr{255, 255, 255, 255}, "255.255.255.255"},
	}
	for _, tt := range hosts {
		actual := tt.input.String()
		if actual != tt.expected {
			t.Errorf("ip stringer(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}
func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()
	if actual != "127.0.0.1" {
		t.Errorf("main expected 127.0.0.1, got %v", actual)
	}
}
