package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	//Given
	r := require.New(t)
	var buf bytes.Buffer
	out = &buf

	// When
	main()

	// Then
	expected := "127.0.0.1\n"
	actual := buf.String()
	r.Equalf(expected, actual, "Unexpected output in main()")
}

var testData = []struct {
	in  IPAddr
	out string
}{
	{IPAddr{127, 0, 0, 1}, "127.0.0.1"},
	{IPAddr{0, 0, 0, 0}, "0.0.0.0"},
	{IPAddr{255, 255, 255, 255}, "255.255.255.255"},
	{IPAddr{1}, "1.0.0.0"},
}

func TestIPAddrFmtInterfaceImpl(t *testing.T) {
	r := require.New(t)
	for _, tt := range testData {
		r.Equalf(tt.out, tt.in.String(), "Convert IP address to dotted quad fails")
	}
}
