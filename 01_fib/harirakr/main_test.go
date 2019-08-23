package main

import (
	"bytes"
	"fmt"
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
	assert.Equal(t, "[1 1 2 3 5 8 13]\n", buf.String())
}

// Test input out of bounds
func TestFibOutOfBounds(t *testing.T) {
	for _, val := range []int64{-94, -93, 93, 94} {
		val := val
		t.Run(fmt.Sprintf("val%v", val), func(t *testing.T) {
			var buf bytes.Buffer
			out = &buf
			fib(val)
			assert.Equal(t,
				"fib: Range out of bounds; takes input > -93 and < 93\n",
				buf.String())
		})
	}
}
