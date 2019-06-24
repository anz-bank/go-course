package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert := assert.New(t)
	testCases := []struct {
		number   int
		expected []int
	}{
		{7, []int{1, 1, 2, 3, 5, 8, 13}},
		{1, []int{1}},
		{2, []int{1, 1}},
		{-6, []int{}},
	}
	for _, testCase := range testCases {
		slice := fibonacci(testCase.number)
		assert.Equal(testCase.expected, slice)
	}
}

func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "1,1,2,3,5,8,13,"
	actual := buf.String()
	actual = strings.Replace(actual, "\n", ",", -1)
	assert.Equal(expected, actual)
}
