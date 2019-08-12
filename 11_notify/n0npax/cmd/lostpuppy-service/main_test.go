package main

import (
	"errors"
	"os"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	os.Args = []string{"", "-p", "88888888"}

	fakeExit := func(int) {
		panic("foo-arg")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()
	assert.Panics(t, main)
}

func TestMainFakeArgs(t *testing.T) {
	os.Args = []string{"", "-s", "map", "-d", "/dev/null", "-p", "8888"}

	parser = func([]string) (int, error) {
		return 80, errors.New("test")
	}
	logFatalf = func(...interface{}) {
		panic("test")
	}
	assert.Panics(t, main)
}

func TestParseArgs(t *testing.T) {
	args := []string{"--port", "1234"}
	port, err := parseArgs(args)
	assert.NoError(t, err)
	assert.Equal(t, port, 1234)
}
