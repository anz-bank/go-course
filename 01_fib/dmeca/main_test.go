package main

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestFibonacci(t *testing.T) {

	testCases := map[string]struct {
		input int
		correctValue string
	}{
		"positive": {input: 7, correctValue: "1\n1\n2\n3\n5\n8\n13\n"},
		"negative": {input: -7, correctValue: "1\n-1\n2\n-3\n5\n-8\n13\n"},
		"zero":     {input: 0, correctValue: ""},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			testResult := captureStdout(func() {
				fib(test.input)
			})
			if testResult != test.correctValue {
				t.Errorf("expected %v, got %v", test.correctValue, testResult)
			}
		})
	}
}

// !!!Not sure how to read output from command line... found this routine from http://craigwickesser.com/2015/01/capture-stdout-in-go/
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
