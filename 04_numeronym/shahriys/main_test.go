package main

import (
	"bytes"
	"fmt"
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
func TestNumeronym(t *testing.T) {

	type test struct {
		expected []string
		actual   []string
	}

	tests := []test{
		{expected: []string{"a2c"}, actual: numeronyms("abac")},
		{expected: []string{}, actual: numeronyms("")},
		{expected: []string{"a"}, actual: numeronyms("a")},
		{expected: []string{"aa", "s3s", "s10g", ""}, actual: numeronyms("aa", "sssss", "sdhjhjhjhjfg", "")},
	}

	for _, tc := range tests {
		if len(tc.expected) > 0 && len(tc.actual) > 0 {
			if !reflect.DeepEqual(tc.expected, tc.actual) {
				fmt.Printf("%v  %T", tc.expected, tc.expected)
				fmt.Printf("%v  %T", tc.actual, tc.actual)
				t.Fatalf(" expected: %v, got: %v", tc.expected, tc.actual)
			}
		}
	}

}
