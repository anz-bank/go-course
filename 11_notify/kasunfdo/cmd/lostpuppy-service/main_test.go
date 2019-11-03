package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Args = []string{""}
	os.Exit(m.Run())
}
