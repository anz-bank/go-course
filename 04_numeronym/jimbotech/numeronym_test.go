package main

import (
	"bytes"
	"reflect"
	"testing"
)

type numeronymTest struct {
	input    string
	expected string
}

var numTests = []numeronymTest{
	{"Kubernetes", "K8s"},
	{"I love You ", "143"},
	{"localizability ", "L12y"},
	{"   Interoperability ", "i14y"},
	{"", ""},
	{"year 2038", "Y2K38"},
	{"somebody said hello", "s17o"},
}

type multiTest struct {
	input    []string
	expected []string
}

var multiNumTests = []multiTest{
	{[]string{"Kubernetes", "localizability", "Interoperability", "I love You ", "abc"},
		[]string{"K8s", "L12y", "i14y", "143", "abc"}},
}

func TestWordMixer(t *testing.T) {

	f := wordMixer()
	for _, tt := range numTests {
		actual := f(tt.input)
		if actual != tt.expected {
			t.Errorf("wordMixer(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}

func TestNumeronyms(t *testing.T) {

	for _, tt := range multiNumTests {
		actual := numeronyms(tt.input...)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("numeronyms(%v): expected %v, actual %v", tt.input, tt.expected, actual)
		}
	}
}

func TestMain(t *testing.T) {
	expected := "[a11y K8s abc]\n"
	var buf bytes.Buffer
	out = &buf
	main()
	actual := buf.String()

	if actual != expected {
		t.Errorf("main expected %v, got %v", expected, actual)
	}
}
