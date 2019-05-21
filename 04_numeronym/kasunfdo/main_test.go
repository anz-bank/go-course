package main

import (
	"bytes"
	"reflect"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote(`[a11y K8s abc]
`)
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestPunctuatedString(t *testing.T) {
	expected := []string{"P23r", "W6t", "G11s"}

	actual := numeronyms("Prozess-Daten-Beschleuniger", "Web Dev'mt", "Game of Thrones")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestAlphaNumericStr(t *testing.T) {
	expected := []string{"W6l"}

	actual := numeronyms("W3school")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}

func TestPunctuatedAlphaNumericStr(t *testing.T) {
	expected := []string{"G6s"}

	actual := numeronyms("Game f' T5s")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Unexpected output in main()\nexpected: %q\nactual: %q", expected, actual)
	}
}
