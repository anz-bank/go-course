package main

import (
	"bytes"
	"strconv"
	"testing"
)

// Test negative Fibonacci series
func TestNegativeFib(t *testing.T) {

	// Test -10 input
	var buf1 bytes.Buffer
	out = &buf1

	fib(-10)
	expected := strconv.Quote("1\n-1\n2\n-3\n5\n-8\n13\n-21\n34\n-55\n")
	actual := strconv.Quote(buf1.String())

	if expected != actual {
		t.Errorf("Test fail fib(-10). Expected=%v, Actual=%v", expected, actual)
	}

	// Test -2 input
	var buf2 bytes.Buffer
	out = &buf2

	fib(-2)
	expected = strconv.Quote("1\n-1\n")
	actual = strconv.Quote(buf2.String())

	if expected != actual {
		t.Errorf("Test fail fib(-2). Expected=%v, Actual=%v", expected, actual)
	}

	// Test -1 input
	var buf3 bytes.Buffer
	out = &buf3

	fib(-1)
	expected = strconv.Quote("1\n")
	actual = strconv.Quote(buf3.String())

	if expected != actual {
		t.Errorf("Test fail fib(-1). Expected=%v, Actual=%v", expected, actual)
	}
}

// Test Fibonacci series
func TestFib(t *testing.T) {

	// Test positive series
	var buf1 bytes.Buffer
	out = &buf1

	fib(10)

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n21\n34\n55\n")
	actual := strconv.Quote(buf1.String())

	if expected != actual {
		t.Errorf("Test fail fib(10). Expected=%v, Actual=%v", expected, actual)
	}

	// Test input 0
	var buf2 bytes.Buffer
	out = &buf2

	fib(0)

	expected = strconv.Quote("")
	actual = strconv.Quote(buf2.String())

	if expected != actual {
		t.Errorf("Test fail fib(0). Expected=%v, Actual=%v", expected, actual)
	}
}

// Test main function
func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf

	main()

	expected := strconv.Quote("1\n1\n2\n3\n5\n8\n13\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Test fail main(). Expected=%v, Actual=%v", expected, actual)
	}

}

// Test input out of bounds
func TestFibOutOfBounds(t *testing.T) {
	var buf1 bytes.Buffer
	out = &buf1

	fib(93)
	expected := strconv.Quote("fib: Range out of bounds; takes input > -93 and < 93\n")
	actual := strconv.Quote(buf1.String())

	if expected != actual {
		t.Errorf("Test fail fib(93). Expected=%v, Actual=%v", expected, actual)
	}

	var buf2 bytes.Buffer
	out = &buf2

	fib(-93)
	expected = strconv.Quote("fib: Range out of bounds; takes input > -93 and < 93\n")
	actual = strconv.Quote(buf2.String())

	if expected != actual {
		t.Errorf("Test fail fib(-93). Expected=%v, Actual=%v", expected, actual)
	}

	var buf3 bytes.Buffer
	out = &buf3

	fib(94)
	expected = strconv.Quote("fib: Range out of bounds; takes input > -93 and < 93\n")
	actual = strconv.Quote(buf3.String())

	if expected != actual {
		t.Errorf("Test fail fib(94). Expected=%v, Actual=%v", expected, actual)
	}

	var buf4 bytes.Buffer
	out = &buf4

	fib(-94)
	expected = strconv.Quote("fib: Range out of bounds; takes input > -93 and < 93\n")
	actual = strconv.Quote(buf4.String())

	if expected != actual {
		t.Errorf("Test fail fib(-94). Expected=%v, Actual=%v", expected, actual)
	}
}
