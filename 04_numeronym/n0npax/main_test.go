package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("[K8s anz i18n]\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected, actual)
	}
}

func TestNumeronyms(t *testing.T) {
	actual := numeronyms("K8bernetes", "anz", "internationalization")
	expected := []string{"K8s", "anz", "i18n"}
	for i, v := range actual {
		if expected[i] != v {
			t.Errorf("Unexpected output. Expected: %q - Actual: %q", expected[i], v)
		}
	}
}

func TestShorten(t *testing.T) {
	testCases := map[string]struct {
		input    string
		expected string
	}{
		"anz":                  {input: "anz", expected: "anz"},
		"Kubernetes":           {input: "Kubernetes", expected: "K8s"},
		"internationalization": {input: "internationalization", expected: "i18n"},
	}
	for name, test := range testCases {
		t.Run(name, helperTestShorten(test.input, test.expected))
	}
}

func helperTestShorten(input string, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := shorten(input)
		if actual != expected {
			t.Errorf("Expected: %q - Actual: %q", expected, actual)
		}
	}
}
