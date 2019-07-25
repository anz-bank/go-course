package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPAddrString(t *testing.T) {
	assert.Equal(t, "0.0.0.0", IPAddr{}.String())
	assert.Equal(t, "1.2.0.0", IPAddr{ip1: 1, ip2: 2}.String())
	assert.Equal(t, "1.2.3.4", IPAddr{1, 2, 3, 4}.String())
	assert.Equal(t, "255.255.255.255", IPAddr{255, 255, 255, 255}.String())
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Equal(t, "127.0.0.1\n", buf.String())
}
