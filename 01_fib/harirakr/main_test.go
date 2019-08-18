package main

import (
	"bytes"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNegativeFib(t *testing.T) {
	assert.Equal(t, []int64{1, -1, 2, -3, 5, -8, 13, -21, 34, -55}, fib(-10))
	assert.Equal(t, []int64{1, -1, 2}, fib(-3))
	assert.Equal(t, []int64{1, -1}, fib(-2))
	assert.Equal(t, []int64{1}, fib(-1))
	assert.Equal(t, []int64{}, fib(0))
}

func TestFib(t *testing.T) {
	assert.Equal(t,
		[]int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765},
		fib(20))
	assert.Equal(t, []int64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}, fib(10))
	assert.Equal(t, []int64{1, 1}, fib(2))
	assert.Equal(t, []int64{1}, fib(1))
	assert.Equal(t, []int64{}, fib(0))
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	expected := strconv.Quote("[1 1 2 3 5 8 13]\n")
	actual := strconv.Quote(buf.String())
	if expected != actual {
		t.Errorf("main fail; expected=%v, actual=%v", expected, actual)
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
