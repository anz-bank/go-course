package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	expected := "retrieved: kelpie brown indispensable\n"
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()
	if actual != expected {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}
