package main

import (
	"bytes"
	"testing"
)

func TestMainOutput(t *testing.T) {
	var buf bytes.Buffer
	mainout = &buf

	main()

	expected := `Fibonacci series starting...
Fibonacci series completed...
`
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in main()")
	}
}

func TestFibOutputFibZero(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(0)

	expected := ``
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in fib(0)")
	}
}

func TestFibOutputFibSeven(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(7)

	expected := `1
1
2
3
5
8
13
`
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in fib(7)")
	}
}

func TestFibOutputFibNine(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(9)

	expected := `1
1
2
3
5
8
13
21
34
`
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in fib(9)")
	}
}

func TestFibOutputNegaFibNine(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(-9)

	expected := `1
-1
2
-3
5
-8
13
-21
34
`
	actual := buf.String()
	if expected != actual {
		t.Errorf("Unexpected output in fib(-9)")
	}
}

func TestFibOutputOutsideRange(t *testing.T) {
	var buf bytes.Buffer
	fibout = &buf

	fib(93)

	expected := `Value outside allowable range (-92 to 92)
`
	actual := buf.String()
	if expected != actual {
		t.Errorf(expected, actual, "Unexpected output in fib(-93)")
	}
}
