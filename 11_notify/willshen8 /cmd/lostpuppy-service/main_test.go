package main

import (
	"os"
	"testing"
)

// main() only starts the server, therefore test case only ensure it runs smoothly
func TestMain(m *testing.M) {
	runMain()
	os.Exit(m.Run())
}
