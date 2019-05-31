package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextVal(t *testing.T) {

	var getNewPair = nextVal(0, 1, 2, 0)

	var actual = getNewPair[0]
	var expected = 1
	if expected != actual {
		t.Errorf("expected %v, got %v", actual, expected)
		t.Fail()
	}
}

func TestFib(t *testing.T) {
	a := assert.New(t)
	var getArray = fib(4)
	a.Contains(getArray, 0, "Array should contain 0")
	a.Contains(getArray, 1, "Array should contain 1")
	a.Contains(getArray, 2, "Array should contain 2")
}
