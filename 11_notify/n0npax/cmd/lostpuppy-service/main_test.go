package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "-p", "--", "-8080"}
	logFatalf = func(...interface{}) {
		panic("test main")
	}
	assert.Panics(t, main, "test main")
}

func TestMainFakeArgs(t *testing.T) {
	os.Args = []string{"", "-s", "map", "-d", "/dev/null", "-p", "8888"}

	parser = func([]string) (int, error) {
		return 80, errors.New("test main fake args")
	}
	logFatalf = func(...interface{}) {
		panic("test main fake args")
	}
	assert.Panics(t, main)
}

func TestParseArgs(t *testing.T) {
	args := []string{"--port", "1234"}
	port, err := parseArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, port, 1234)
}
