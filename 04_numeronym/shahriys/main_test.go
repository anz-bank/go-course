package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	var bufnum bytes.Buffer
	outnum = &bufnum

	main()
	expected := `[a11y K8s abc]`
	actual := bufnum.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()" + expected + " " + actual)
	}

}
func TestLettersSingle(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{"Test1": {"abac", []string{"a2c"}},
		"Test2": {"", []string{""}},
		"Test3": {"a", []string{"a"}},
	}

	for name, test := range cases {
		actual := numeronyms(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(" test:%v input:%v, expected: %v, got: %v", name, test.input, test.expected, actual)
		}
	}

}
func TestLettersMulti(t *testing.T) {
	cases := map[string]struct {
		input    []string
		expected []string
	}{
		"Test4": {nil, []string{}},
		"Test5": {[]string{"aa", "sssss", "sdhjhjhjhjfg", ""}, []string{"aa", "s3s", "s10g", ""}},
		"Test6": {[]string{}, []string{}},
	}

	for name, test := range cases {
		actual := numeronyms(test.input...)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf(" test:%v input:%v, expected: %v, got: %v", name, test.input, test.expected, actual)
		}
	}

}
