package main

import (
	"bytes"
	"strconv"
	"testing"
)

func TestMainCallsFib(t *testing.T) {
	// Arrange
	var buf bytes.Buffer
	out = &buf

	// Act
	main()

	// Assert
	expected := strconv.Quote("0\n1\n1\n2\n3\n5\n8\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("Unexpected behavior with test input of 7")
		t.Errorf("expected: %s\nactual:%s", expected, actual)
	}

}

func TestFibInputPositive(t *testing.T) {
	// Arrange
	var buf bytes.Buffer
	out = &buf

	// Act
	fib(7)

	// Assert
	expected := strconv.Quote("0\n1\n1\n2\n3\n5\n8\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected behavior with test input of 7")
		t.Errorf("expected: %s\nactual:%s", expected, actual)
	}
}

func TestFibInputNegative(t *testing.T) {
	// Arrange
	var buf bytes.Buffer
	out = &buf

	// Act
	fib(-7)

	// Assert
	expected := strconv.Quote("0\n1\n-1\n2\n-3\n5\n-8\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected output in main()")
		t.Errorf("expected: " + expected + "\nactual: " + actual)
	}
}

func TestFibInputZero(t *testing.T) {
	// Arrange
	var buf bytes.Buffer
	out = &buf

	// Act
	fib(0)

	// Assert
	expected := strconv.Quote("0\n")
	actual := strconv.Quote(buf.String())

	if expected != actual {
		t.Errorf("unexpected output in main()")
		t.Errorf("expected: " + expected + "\nactual: " + actual)
	}
}

func TestCalcPositiveFib(t *testing.T) {
	var fibResult = calcPositiveFib(4)
	var expectedResult int64 = 3
	if fibResult != expectedResult {
		t.Errorf("wrong output from calc fib")
		t.Errorf("expected: %d\nactual:%d", expectedResult, fibResult)
	}
}

func TestCalcNegativeFib(t *testing.T) {
	var fibResult = calcNegativeFib(-4)
	var expectedResult int64 = -3
	if fibResult != expectedResult {
		t.Errorf("wrong output from calc fib")
		t.Errorf("expected: %d\nactual:%d", expectedResult, fibResult)
	}
}

func TestCalcZeroFib(t *testing.T) {
	var expectedResult int64

	var fibResult = calcPositiveFib(0)
	if fibResult != expectedResult {
		t.Errorf("wrong output from calc fib")
		t.Errorf("expected: %d\nactual:%d", expectedResult, fibResult)
	}
}
