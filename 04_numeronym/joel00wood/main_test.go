package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

// enter borrowed code
// https://gist.github.com/hauxe/e935a7f9012bf2649710cf75af323dbf
// stdOut capture helper
func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, err := io.Copy(&buf, reader)
		if err != nil {
			panic(err)
		}
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func TestMain(t *testing.T) {
	expected := "[a11y K8s abc]\n"
	actual := captureOutput(func() { main() })
	if expected != actual {
		t.Errorf("Unexpected result in main(), expected=%q, got=%q", expected, actual)
	}
}

var testCases = map[string]struct {
	input    []string
	expected []string
}{
	"Standard input": {
		[]string{"accessibility", "Kubernetes", "abc"},
		[]string{"a11y", "K8s", "abc"},
	},
	"Empty input": {
		[]string{},
		[]string{},
	},
	"Single Double Triple Quad": {
		[]string{"a", "ab", "abc", "abcd"},
		[]string{"a", "ab", "abc", "a2d"},
	},
	"Some more complicated words": {
		[]string{"feremone", "guttulous", "synchronous", "isopropenyl", "mericarp",
			"spleened", "Universalian", "appeacher", "antidisestablishmentarianism"},
		[]string{"f6e", "g7s", "s9s", "i9l", "m6p", "s6d", "U10n", "a7r", "a26m"},
	},
}

func TestNumeronym(t *testing.T) {
	for name, test := range testCases {
		for i := 0; i < len(test.input); i++ {
			actual := numeronym(test.input[i])
			expected := test.expected[i]
			if actual != expected {
				t.Errorf("Unexpected result for %q numeronym(%q), expected=%q, got=%q",
					name, test.input[i], test.expected[i], actual)
			}
		}
	}
}

func TestNumeronyms(t *testing.T) {
	for name, test := range testCases {
		actual := numeronyms(test.input...)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("Unexpected result for %q numeronyms(%v), expected=%v, got=%v",
				name, test.input, test.expected, actual)
		}
	}
}
