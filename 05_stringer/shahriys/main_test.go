package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	var bufstringer bytes.Buffer
	outstringer = &bufstringer

	main()
	expected := `127.0.0.1`
	actual := bufstringer.String()

	if expected != actual {
		t.Errorf("Unexpected output in main()" + expected + " " + actual)
	}

}
func TestIpAddr(t *testing.T) {

	type test struct {
		input    IPAddr
		expected string
	}

	tests := []test{
		{input: IPAddr{127, 0, 0, 1}, expected: `127.0.0.1`},
		{input: IPAddr{}, expected: `0.0.0.0`},
	}

	for _, tc := range tests {
		a := tc.input
		actual := a.String()
		if !reflect.DeepEqual(tc.expected, actual) {
			fmt.Printf("%v  %T", tc.expected, tc.expected)
			fmt.Printf("%v  %T", actual, actual)
			t.Errorf(" input:%v, expected: %v, got: %v", tc.input, tc.expected, actual)
		}

	}

}
