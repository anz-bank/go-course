package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetters(t *testing.T) {
	testData := []struct {
		Scenario string
		input    string
		expected map[rune]int
	}{
		{"Scenario One", "eaa!baabq23123!", map[int32]int{33: 2, 49: 1, 50: 2, 51: 2, 97: 4, 98: 2, 101: 1, 113: 1}},
		{"Scenario two", "", map[int32]int{}},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			letterMap := letters(td.input)
			assert.Equal(t, td.expected, letterMap)
		})
	}
}
func TestMain(t *testing.T) {
	assert := assert.New(t)
	var buf bytes.Buffer
	out = &buf
	main()
	expected := "a:2,b:1"
	actual := buf.String()
	actual = strings.Replace(actual, "\n", ",", -1)
	actual = strings.TrimRight(actual, ",")
	assert.Equal(expected, actual)
}

func TestSortLetters(t *testing.T) {
	testData := []struct {
		Scenario string
		input    map[rune]int
		expected []string
	}{
		{"Scenario One", map[int32]int{33: 2, 49: 1, 50: 2, 51: 2, 97: 4, 98: 2, 101: 1, 113: 1},
			[]string{"!:2", "1:1", "2:2", "3:2", "a:4", "b:2", "e:1", "q:1"}},
		{"Scenario two", map[int32]int{}, []string{}},
	}
	for _, td := range testData {
		td := td
		t.Run(td.Scenario, func(t *testing.T) {
			sortedLetter := sortLetters(td.input)
			assert.Equal(t, td.expected, sortedLetter)
		})
	}
}
