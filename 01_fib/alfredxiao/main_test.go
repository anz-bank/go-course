package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibEdgeCases(t *testing.T) {
	assert.Equal(t, []int{0}, fibSeries(0))
	assert.Equal(t, []int{0, 1}, fibSeries(1))
	assert.Equal(t, []int{0, 1}, fibSeries(-1))
	assert.Equal(t, []int{0, 1, 1}, fibSeries(2))
	assert.Equal(t, []int{0, 1, -1}, fibSeries(-2))
}

func TestFibPositive(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 3, 5, 8, 13}, fibSeries(7))
}

func TestFibNegative(t *testing.T) {
	assert.Equal(t, []int{0, 1, -1, 2, -3, 5, -8, 13}, fibSeries(-7))
}

func TestFibSeriesPanic(t *testing.T) {
	assert.Panics(t, func() { fibSeries(93) }, "Panic when number of range")
	assert.Panics(t, func() { fibSeries(94) }, "Panic when number of range")
	assert.Panics(t, func() { fibSeries(-93) }, "Panic when number of range")
	assert.Panics(t, func() { fibSeries(-94) }, "Panic when number of range")
}

func TestFib(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	fib(2)
	assert.Equal(t, "0\n1\n1\n", buf.String())
}

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	out = &buf
	main()
	assert.Contains(t, buf.String(), "3\n5\n8")
}
