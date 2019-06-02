package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIpAddress(t *testing.T) {
	var expected string
	var actual string
	expected = "127.0.0.1"
	actual = IPAddr{127, 0, 0, 1}.String()
	if expected != actual {
		t.Errorf("expected %v, got %v", actual, expected)
		t.Fail()
	}
}

func TestMain(m *testing.M) {
	main()
	fmt.Printf("test!")
	os.Exit(m.Run())
}
