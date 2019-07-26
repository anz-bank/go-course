package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	testData := []struct {
		Scenario string
		unSorted []int
		expected []int
	}{
		{"Happy Day Scenario", []int{2, 5, 8, 1, 13, 1, 3}, []int{1, 1, 2, 3, 5, 8, 13}},
		{"Slice containing negative numbers", []int{-1, 3, 5, -7}, []int{-7, -1, 3, 5}},
		{"Empty Slice", []int{}, []int{}},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			sorted := bubble(td.unSorted)
			assert.Equal(t, td.expected, sorted)
		})
	}
}
func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "[1 2 3 5]"
	actual := buf.String()
	assert.Equal(expected, actual)
}
